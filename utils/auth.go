package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// Auth middleware
func RequireAuth(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	envMap, err := godotenv.Read(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			token, err := jwt.Parse(r.Header["Authorization"][1], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized!"))
				}
				return envMap["JWT_SECRET"], nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized!" + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized!"))
		}
	})
}

func GenerateToken() (string, error) {
	envMap, err := godotenv.Read(".env")

	if err != nil {
		return "Error loading .env file", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix()
	tokenStr, err := token.SignedString([]byte(envMap["JWT_SECRET"]))
	return tokenStr, nil
}
