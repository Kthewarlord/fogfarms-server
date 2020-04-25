package jwt

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"

	"github.com/KitaPDev/fogfarms-server/src/util/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/gddo/httputil/header"
	"github.com/gorilla/securecookie"
)

const (
	bearer       string = "bearer"
	bearerFormat string = "Bearer "
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//var jwtKey := os.Getenv("SECRET_KEY_JWT")
var jwtKey = "s"

var secureCookie = securecookie.New([]byte(jwtKey), nil)

func AuthenticateUserToken(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("jwtToken")
	if err != nil {
		if err == http.ErrNoCookie {
			msg := "Error: No Token Found"
			http.Error(w, msg, http.StatusUnauthorized)
			log.Println(err)
			return false
		}

		msg := `Error: Failed to Retrieve Token from Cookie"`
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return false
	}

	var tokenString string
	//tokenString := cookie.Value
	err = secureCookie.Decode("jwtToken", cookie.Value, &tokenString)
	if err != nil {
		msg := `Error: Failed to Decode Token Value"`
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return false
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			msg := "Error: Invalid Signature"
			http.Error(w, msg, http.StatusUnauthorized)
			log.Println(err)
			return false
		}

		msg := "Error: Failed to Parse Token"
		http.Error(w, msg, http.StatusUnauthorized)
		http.Error(w, msg, http.StatusUnauthorized)
		log.Println(err)
		return false
	}

	if !token.Valid {
		msg := "Error: Invalid Token"
		http.Error(w, msg, http.StatusUnauthorized)
		return false
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 0*time.Second {

		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 10*time.Minute {
			InvalidateToken(w)
			log.Println("Token exceeded timeout")
			msg := "Error: Token Exceeded Timeout Limit, Sign In Again"
			http.Error(w, msg, http.StatusUnauthorized)
			return false

		} else {
			log.Println("UserID: ", claims.Username)
			GenerateToken(claims.Username, w)
		}

	}

	return true
}

func AuthenticateSignIn(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	var credentials Input

	success := jsonhandler.DecodeJsonFromBody(w, r, &credentials)
	if !success {
		return
	}
	log.Println(credentials, "  ... ", &credentials)
	username := credentials.Username
	password := credentials.Password

	exists, _, err := user.ExistsByUsername(username)
	if err != nil {
		msg := "Error: Failed to Exists By Username"
		http.Error(w, msg, http.StatusUnauthorized)
		log.Println(err)
		return
	}
	if !exists {
		msg := "Error: User Not Found"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	valid, err := user.AuthenticateByUsername(username, password)
	if err != nil {
		msg := "Error: Failed to Authenticate By UserID"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if !valid {
		msg := "Error: Failed to Invalid Credentials"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	err = GenerateToken(username, w)
	if err != nil {
		msg := "Error: Failed to Generate Token"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful"))
}

func GenerateToken(username string, w http.ResponseWriter) error {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		msg := "Error: Failed to Generate Token"
		http.Error(w, msg, http.StatusUnauthorized)
		return err
	}
	encoded, err := secureCookie.Encode("jwtToken", tokenString)
	if err != nil {
		msg := "Error: Failed to Encode Cookie"
		http.Error(w, msg, http.StatusInternalServerError)
		log.Println(err)
		return err
	}

	cookie := &http.Cookie{
		Name:    "jwtToken",
		Value:   encoded,
		Expires: expirationTime,
		Path:    "/",
	}
	http.SetCookie(w, cookie)
	cs := w.Header().Get("Set-Cookie")
	//	cs += "; SameSite=None; Secure;"
	w.Header().Set("Set-Cookie", cs)
	return nil
}

func InvalidateToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwtToken",
		Value:   "",
		Expires: time.Now(),
		Path:    "/",
	})
}

func extractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], bearer) {
		return "", false
	}

	return authHeaderParts[1], true
}

func generateAuthHeaderFromToken(token string) string {
	return bearerFormat + token
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		NewPassword string `json:"new_password"`
	}
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	var credentials Input

	success := jsonhandler.DecodeJsonFromBody(w, r, &credentials)
	if !success {
		return
	}
	username := credentials.Username
	password := credentials.Password
	valid, err := user.AuthenticateByUsername(username, password)
	if err != nil {
		msg := "Error: Failed to Authenticate By UserID"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	if !valid {
		msg := "Error: Failed to Change password"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}
	err = user.ChangePassword(credentials.Username, credentials.NewPassword)
	if err != nil {
		msg := "Error: Failed to Change password"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

}
