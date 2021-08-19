package migration

import "time"

func (User) TableName() string {
	return "master_user"
}

type User struct {
	Id         int64     `json:"id" gorm:"column:id;primary_key"`
	Username   *string   `json:"username" gorm:"column:username"`
	Password   *string   `json:"password" gorm:"column:password"`
	Email      *string   `json:"email" gorm:"column:email"`
	Isdeleted  bool      `json:"is_deleted" gorm:"column:is_deleted"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at" `
	ModifiedAt time.Time `json:"modified_at" gorm:"column:modified_at" `
}
