package token

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/putranurf/be-tpd/model/gorm"
	migration "github.com/putranurf/be-tpd/model/migration/token"
)

func CreateToken(c echo.Context) error {

	name := c.FormValue("name")
	secret_key := c.FormValue("secret_key")
	device_id := c.FormValue("device_id")
	device_type := c.FormValue("device_type")
	// Set custom claims
	claims := &migration.JwtCustomClaims{
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
		time.Now(),
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))
	//Set Array Object For API
	arrobj := &migration.Token{
		0,
		name,
		device_id,
		device_type,
		t,
		time.Unix(time.Now().Add(time.Minute*15).Unix(), 0),
		time.Now(),
		secret_key,
	}
	result, err := gorm.CreateToken(arrobj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  echo.ErrUnauthorized.Code,
			"message": echo.ErrUnauthorized.Message,
		})
	}
	//Set Token Header
	c.Response().Header().Set("token", t)
	return c.JSON(http.StatusOK, result)

}
