package token

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/putranurf/be-tpd/model/gorm"
	migration "github.com/putranurf/be-tpd/model/migration/token"
)

type M map[string]interface{}

//Backup Fetch Token
// func FetchToken(c echo.Context) error {

// 	name := c.FormValue("name")
// 	secret_key := c.FormValue("secret_key")
// 	device_id := c.FormValue("device_id")
// 	device_type := c.FormValue("device_type")

// 	// Set custom claims
// 	claims := &migration.JwtCustomClaims{
// 		name,
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
// 		},
// 		time.Now(),
// 	}

// 	// Create token with claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Generate encoded token and send it as response.
// 	t, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

// 	//Set Array Object For API
// 	arrobj := &migration.Token{
// 		0,
// 		name,
// 		device_id,
// 		device_type,
// 		t,
// 		time.Unix(time.Now().Add(time.Minute*30).Unix(), 0),
// 		time.Now(),
// 		secret_key,
// 	}

// 	result, err := gorm.FetchToken(arrobj)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"status":  echo.ErrUnauthorized.Code,
// 			"message": echo.ErrUnauthorized.Message,
// 		})
// 	}

// 	//Set Token Header
// 	c.Response().Header().Set("token", t)
// 	return c.JSON(http.StatusOK, result)

// }

func FetchToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}
	fmt.Print(r.Method)
	name := r.FormValue("name")
	secret_key := r.FormValue("secret_key")
	device_id := r.FormValue("device_id")
	device_type := r.FormValue("device_type")

	claims := &migration.JwtCustomClaims{
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		time.Now(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte("the secret of kalimdor"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Set Array Object For API
	arrobj := &migration.Token{
		0,
		name,
		device_id,
		device_type,
		signedToken,
		time.Unix(time.Now().Add(time.Minute*1).Unix(), 0),
		time.Now(),
		secret_key,
	}

	result, err := gorm.FetchToken(arrobj)
	if err != nil {
		return
	}

	tokenString, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenString)

}
