package gorm

import (
	"fmt"
	"net/http"
	"time"

	"github.com/putranurf/be-tpd/db"
	password "github.com/putranurf/be-tpd/internal/repo/migration/auth"
	migration "github.com/putranurf/be-tpd/internal/repo/migration/user"
	"github.com/putranurf/be-tpd/pkg/helpers"
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

func CreateReset(ctx interface{}) (Response, error) {
	var res Response
	origin := ctx.(*password.Password)
	arrobj := password.Password{
		Id:          origin.Id,
		Newpassword: origin.Newpassword,
	}
	if err := db.GetDBInstance().Model(&migration.User{}).Where("id = ?", arrobj.Id).Select("password", "modified_at").
		Updates(map[string]interface{}{
			"password": arrobj.Newpassword, "modified_at": time.Now(),
		}).Error; err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = nil
	return res, nil
}
