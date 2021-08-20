package gorm

import (
	"fmt"
	"net/http"

	"github.com/putranurf/be-tpd/db"
	"github.com/putranurf/be-tpd/helpers"
	migration "github.com/putranurf/be-tpd/repo/migration/user"
)

func GetLogin(username, password string) (Response, error) {
	var obj migration.User
	var arrobj []migration.User
	var res Response
	if db := db.GetDBInstance().Where("username = ?", username).First(&obj); db.Error != nil {
		fmt.Println("Username Doesn't Match")
		return res, db.Error
	}
	match, err := helpers.CheckPasswordHash(password, *obj.Password)
	if !match {
		fmt.Println("Password Doesn't Match")
		return res, err
	}
	arrobj = append(arrobj, obj)
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj
	return res, nil
}

// func CreateForgot(email string) (Response, error) {
// }
