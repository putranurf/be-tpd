package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("masuk middleware")
		if r.URL.Path == "/api/create-token" {
			next.ServeHTTP(w, r)
			return
		}
		//Check Token Valid
		authorizationHeader := r.Header.Get("token")
		if !strings.Contains(authorizationHeader, "") {
			// http.Error(w, "Invalid token", http.StatusBadRequest)
			err, _ := json.Marshal(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Invalid Token",
			})
			w.Header().Set("Content-Type", "application/json")
			w.Write(err)
			return
		}
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
		})
		//Check Token Exists
		if err != nil {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			err, _ := json.Marshal(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			w.Header().Set("Content-Type", "application/json")
			w.Write(err)
			return
		}
		//Check Token Expired
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			err, _ := json.Marshal(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			w.Header().Set("Content-Type", "application/json")
			w.Write(err)
			return
		}
		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
