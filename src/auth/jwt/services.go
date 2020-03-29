package jwt

import (
	"../../user"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	bearer       string = "bearer"
	bearerFormat string = "Bearer %s"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) bool {
	jwtKey := os.Getenv("SECRET_KEY_JWT")

	cookie, err := r.Cookie("jwtToken")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}

		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	tokenString := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	return true
}

func AuthenticateSignIn(w http.ResponseWriter, r *http.Request) {
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	exists, _ := user.Exists(username)
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		log.Fatal(io.WriteString(w, `{"error":"user_not_found"}"`))
		return
	}

	valid := user.ValidateUser(username, password)
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	GenerateToken(username, w)
}

func GenerateToken(username string, w http.ResponseWriter) {
	jwtKey := os.Getenv("SECRET_KEY_JWT")

	expirationTime := time.Now().Add(10 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp": expirationTime,
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(io.WriteString(w, `{"error":"token_generation_failed"`))
		return
	}

	http.SetCookie(w, &http.Cookie {
		Name: "jwtToken",
		Value: tokenString,
		Expires: expirationTime,
	})
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	jwtKey := os.Getenv("SECRET_KEY_JWT")

	cookie, err := r.Cookie("jwtToken")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	GenerateToken(claims.Username, w)
}

func InvalidateToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name: "jwtToken",
		Value: "",
		Expires: time.Unix(0,0),
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
	return fmt.Sprintf(bearerFormat, token)
}