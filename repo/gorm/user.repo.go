package gorm

import (
	"net/http"
	"time"

	"github.com/putranurf/be-tpd/db"
	migration "github.com/putranurf/be-tpd/repo/migration/user"
)

func CreateUser(username, password, email string) (Response, error) {
	var res Response
	arrobj := migration.User{
		Username:  &username,
		Password:  &password,
		Email:     &email,
		Isdeleted: false,
		CreatedAt: time.Now(),
	}
	if err := db.GetDBInstance().Select("Username", "Password", "Email", "Isdeleted", "CreatedAt").Create(&arrobj).Error; err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = nil
	return res, nil
}

func ListUser() (Response, error) {
	var arrobj []migration.User
	var res Response
	if err := db.GetDBInstance().Where("is_deleted = ?", false).Find(&arrobj).Error; err != nil {
		return res, err
	}
	arrobj = append(arrobj)
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj
	return res, nil
}

func UpdateUser(id int, username, password string) (Response, error) {
	var res Response
	if err := db.GetDBInstance().Model(&migration.User{}).Where("id = ?", id).Update("username", username).Error; err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = nil
	return res, nil
}

func DeleteUser(id int) (Response, error) {
	var res Response
	if err := db.GetDBInstance().Model(&migration.User{}).Where("id = ?", id).Select("is_deleted", "modified_at").Updates(map[string]interface{}{"is_deleted": true, "modified_at": time.Now()}).Error; err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = nil
	return res, nil
}
