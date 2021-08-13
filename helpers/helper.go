package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	BASIC_SCHEMA  string = "Basic "
	BEARER_SCHEMA string = "Bearer "
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

//Get Token Header
// func RequestToken(c echo.Context) (string, error) {
// 	token := c.Response().Header().Get("token")

// 	return token, nil
// }

// // FromRequest extracts the auth token from req.
// func FromRequest(req *http.Request) (string, error) {
// 	// Grab the raw Authoirzation header
// 	authHeader := req.Header.Get("Token")
// 	if authHeader == "" {
// 		return "", errors.New("Token header required")
// 	}

// 	return authHeader, nil

// 	// // Confirm the request is sending Basic Authentication credentials.
// 	// if !strings.HasPrefix(authHeader, BASIC_SCHEMA) {
// 	// 	return "", errors.New("Token requires Basic/Bearer scheme")
// 	// }

// 	// // Get the token from the request header
// 	// // The first six characters are skipped - e.g. "Basic ".
// 	// if strings.HasPrefix(authHeader, BASIC_SCHEMA) {
// 	// 	str, err := base64.StdEncoding.DecodeString(authHeader[len(BASIC_SCHEMA):])
// 	// 	if err != nil {
// 	// 		return "", errors.New("Base64 encoding issue")
// 	// 	}
// 	// 	creds := strings.Split(string(str), ":")
// 	// 	return creds[0], nil
// 	// } else {
// 	// 	return authHeader[len(BEARER_SCHEMA):], nil
// 	// }
// }
