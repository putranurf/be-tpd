package migration

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (Token) TableName() string {
	return "token"
}

type Token struct {
	Id           int64     `json:"id" gorm:"column:id;primary_key"`
	Name         string    `json:"name" gorm:"column:name"`
	DeviceId     string    `json:"device_id" gorm:"column:device_id"`
	DeviceType   string    `json:"device_type" gorm:"column:device_type"`
	Token        string    `json:"token" gorm:"column:token"`
	TokenExpired time.Time `json:"token_expired" gorm:"token_expired"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at" `
	SecretKey    string    `json:"secret_key" gorm:"column:secret_key"`
}

type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
	CreatedAt time.Time `json:"created_at"`
}
