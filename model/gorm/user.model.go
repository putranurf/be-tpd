package gorm

import (
	"net/http"

	"github.com/putranurf/be-tpd/db"
	migration "github.com/putranurf/be-tpd/model/migration/user"
)

func GetUsers() (Response, error) {
	var obj migration.User
	var arrobj []migration.User
	var res Response

	if err := db.GetDBInstance().Find(&obj, 1).Error; err != nil {
		return res, err
	}
	arrobj = append(arrobj, obj)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}
