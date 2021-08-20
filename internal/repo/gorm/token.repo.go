package gorm

import (
	"net/http"

	"github.com/putranurf/be-tpd/db"
	migration "github.com/putranurf/be-tpd/internal/repo/migration/token"
)

//Fetch Token
func CreateToken(obj interface{}) (Response, error) {
	var res Response
	origin := obj.(*migration.Token)
	arrobj := migration.Token{
		Name:         origin.Name,
		DeviceId:     origin.DeviceId,
		DeviceType:   origin.DeviceType,
		Token:        origin.Token,
		TokenExpired: origin.TokenExpired,
		CreatedAt:    origin.CreatedAt,
		SecretKey:    origin.SecretKey,
	}
	if err := db.GetDBInstance().Select("Name", "DeviceId", "DeviceType", "Token", "TokenExpired", "CreatedAt", "SecretKey").Create(&arrobj).Error; err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func GetToken() (Response, error) {

	var arrobj []migration.Token
	var res Response

	if err := db.GetDBInstance().Find(&arrobj).Error; err != nil {
		return res, err
	}

	arrobj = append(arrobj)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil

}
