package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	// "github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"

	// "github.com/labstack/echo/middleware"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/putranurf/be-tpd/controller/token"
)

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func Init() {
	// func Init() *echo.Echo {
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })

	// // Token
	// e.POST("/fetch-token", token.FetchToken)

	// // Auth
	// // e.POST("/login", auth.CheckLogin, middleware.IsAuthenticated)
	// e.GET("/generate-hash/:password", auth.GenerateHashPassword)

	// // CRUD
	// e.GET("/users", user.GetUsers)
	// e.GET("/test-struct", controller.TestStructValidation)
	// e.GET("/test-var", controller.TestVarValidation)

	// adminGroup := e.Group("/v1")
	// config := middleware.JWTConfig{
	// 	// Claims:     &migration.JwtCustomClaims{},
	// 	SigningKey: []byte(os.Getenv("TOKEN_SECRET")),
	// }
	// adminGroup.Use(middleware.JWTWithConfig(config))

	// // // Attach jwt token refresher.
	// // adminGroup.Use(middlewares.TokenRefresherMiddleware)

	// adminGroup.GET("/users", user.GetUsers())

	//MUX Router
	mux := new(CustomMux)
	mux.RegisterMiddleware(MiddlewareJWTAuthorization)

	// e.POST("/fetch-token", echo.WrapHandler(http.HandlerFunc(token.FetchToken)))
	mux.HandleFunc("/fetch-token", token.FetchToken)

	mux.HandleFunc("/index", HandlerIndex)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":1234"

	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()

	// return e
	// return mux
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("masuk middleware")
		if r.URL.Path == "/fetch-token" {
			next.ServeHTTP(w, r)
			return
		}
		// fmt.Println("Udah melewati fetch-token")

		authorizationHeader := r.Header.Get("token")
		fmt.Println(authorizationHeader)
		if !strings.Contains(authorizationHeader, "") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		// // ...
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	message := fmt.Sprintf("hello Sudah Masuk API")
	w.Write([]byte(message))
}
