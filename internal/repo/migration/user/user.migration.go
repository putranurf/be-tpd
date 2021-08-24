package migration

import "time"

func (User) TableName() string {
	return "master_user"
}

type User struct {
	Id           int64     `json:"id" gorm:"column:id"`
	IdMasterUser string    `json:"id_master_user" gorm:"column:id_master_user"`
	Username     *string   `json:"username" gorm:"primaryKey;autoIncrement:false;column:username"`
	Password     *string   `json:"password" gorm:"column:password"`
	Email        *string   `json:"email" gorm:"column:email"`
	Isdeleted    bool      `json:"is_deleted" gorm:"column:is_deleted"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at" `
	ModifiedAt   time.Time `json:"modified_at" gorm:"column:modified_at" `
}
