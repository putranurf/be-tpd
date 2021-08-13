package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey:  []byte(os.Getenv("TOKEN_SECRET")),
	TokenLookup: "header:token",
})

// func CheckToken() {
// 	authToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
// 		if jwt.GetSigningMethod("HS256") != token.method {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("TOKEN_SECRET")), nil
// 	})
// 	if authToken != nil {
// 		if !authToken.Valid || err != nil {
// 			log.Println(err)
// 			err = errors.New("Failed auth : Please provide correct token !")
// 		}
// 	} else {
// 		if claims, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {

// 		}
// 	}
// }

func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Header ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("TOKEN_SECRET")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}

func Pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
