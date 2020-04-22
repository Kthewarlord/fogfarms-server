package user

import (
	"log"
	"net/http"
	"time"

	"github.com/KitaPDev/fogfarms-server/src/jsonhandler"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/util/user/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
)

func AuthenticateByUsername(username string, password string) (bool, error) {
	return repository.ValidateUserByUsername(username, password)
}

func CreateUser(w http.ResponseWriter, r *http.Request) error {
	type Input struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		IsAdministrator bool   `json:"is_administrator"`
	}

	var input Input
	success := jsonhandler.DecodeJsonFromBody(w, r, &input)
	if !success {
		return http.ErrBodyNotAllowed
	}

	err := repository.CreateUser(input.Username, input.Password, input.IsAdministrator)
	if err != nil {
		msg := "Error: Failed to Create User"
		http.Error(w, msg, http.StatusBadRequest)
		log.Println(err)
		return err
	}

	return nil
}

func GetAllUsers() ([]models.User, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return users, err
	}

	return users, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByID(userID int) (*models.User, error) {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsersByID(userIDs []int) ([]models.User, error) {
	users, err := repository.GetUsersByID(userIDs)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByUsernameFromCookie(r *http.Request) (*models.User, error) {
	var jwtKey = "s"
	var secureCookie = securecookie.New([]byte(jwtKey), nil)

	type Claims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	cookie, err := r.Cookie("jwtToken")

	var tokenString string
	err = secureCookie.Decode("jwtToken", cookie.Value, &tokenString)

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
	if err != nil {
		return nil, err
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 0*time.Second && token.Valid {
		return GetUserByUsername(claims.Username)

	} else {
		return GetUserByUsername(claims.Username)
	}

}

func GetUserStringByUsernameFromCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	var jwtKey = "s"
	var secureCookie = securecookie.New([]byte(jwtKey), nil)

	type Claims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}
	cookie, err := r.Cookie("jwtToken")
	if err != nil {
		return "", err
	}

	var tokenString string
	err = secureCookie.Decode("jwtToken", cookie.Value, &tokenString)
	if err != nil {
		return "", err
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

	if err != nil {
		return "", err
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 0*time.Second && token.Valid {

		return claims.Username, err
	} else {
		return claims.Username, err
	}

}

func ExistsByUsername(username string) (bool, *models.User, error) {
	if user, err := GetUserByUsername(username); user != nil && err == nil {
		return true, user, nil
	} else {
		return false, nil, err
	}
}

func ExistsByID(userID int) (bool, *models.User, error) {
	if user, err := GetUserByID(userID); user != nil && err == nil {
		return true, user, nil
	} else {
		return false, nil, err
	}
}

func ChangePassword(username string, newPassword string) error {
	return repository.ChangePassword(username, newPassword)
}
