package middleware

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(os.Getenv("TOKEN_SECRET")),
	// TokenLookup: "header:token",
})

// // TokenRefresherMiddleware middleware, which refreshes JWT tokens if the access token is about to expire.
// func TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// If the user is not authenticated (no user token data in the context), don't do anything.
// 		// user := c.Get("user").(*jwt.Token)
// 		// claims := user.Claims.(*migration.JwtCustomClaims)
// 		// name := claims.Name
// 		fmt.Println(c.Get("user"))
// 		// fmt.Println(name)

// 		if c.Get("user") == nil {
// 			return next(c)
// 		}
// 		// Gets user token from the context.
// 		// u := c.Get("user").(*jwt.Token)

// 		claims := (*&migration.Token{})

// 		// We ensure that a new token is not issued until enough time has elapsed
// 		// In this case, a new token will only be issued if the old token is within
// 		// 15 mins of expiry.
// 		if claims.TokenExpired.Sub(time.Now()) < 15*time.Minute {
// 			// Gets the refresh token from the cookie.
// 			rc, err := c.Cookie("refresh-token")
// 			if err == nil && rc != nil {
// 				// Parses token and checks if it valid.
// 				_, err := jwt.ParseWithClaims(rc.Value, migration.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 					return []byte(os.Getenv("TOKEN_SECRET")), nil
// 				})
// 				if err != nil {
// 					if err == jwt.ErrSignatureInvalid {
// 						// c.Response().Writer.WriteHeader(http.StatusUnauthorized)
// 						fmt.Println("error unathorized")
// 					}
// 				}

// 				// if tkn != nil && tkn.Valid {
// 				// 	// If everything is good, update tokens.

// 				// 	_ = token.GenerateTokensAndSetCookies(&user.User{
// 				// 		Name: claims.Name,
// 				// 	}, c)
// 				// }
// 			}
// 		}

// 		return next(c)
// 	}
// }
