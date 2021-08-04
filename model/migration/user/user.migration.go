package migration

func (User) TableName() string {
	return "master_user"
}

type User struct {
	Id       int64   `json:"id" gorm:"column:id;primary_key"`
	Username *string `json:"username" gorm:"column:username"`
	Email    *string `json:"email" gorm:"column:email"`
}
