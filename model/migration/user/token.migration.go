package migration

import "time"

func (Token) TableName() string {
	return "token"
}

type Token struct {
	Id        int64     `json:"id" gorm:"column:id;primary_key"`
	Token     *string   `json:"token" gorm:"column:token"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" `
}
