package jwt

import (
	"../../../models"
	"../../user"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	bearer       string = "bearer"
	bearerFormat string = "Bearer %s"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("SECRET_KEY_JWT")

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if !user.Exists(username) {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"error":"invalid_credentials"}"`)

	} else {
		if
	}

	expectedPassword, ok =

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