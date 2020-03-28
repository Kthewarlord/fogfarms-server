package jwt

import (
	"../../../models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

const (
	bearer       string = "bearer"
	bearerFormat string = "Bearer %s"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateJwtToken(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("SECRET_KEY_JWT")

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.Write(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "jwtToken",
		Value: tokenString
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