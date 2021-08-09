package gorm

import (
	"net/http"
	"time"

	"github.com/putranurf/be-tpd/db"
	migration "github.com/putranurf/be-tpd/model/migration/user"
)

//Fetch Token
func FetchToken(token string) (Response, error) {
	var res Response
	arrobj := migration.Token{
		Token:     &token,
		CreatedAt: time.Now(),
	}

	if err := db.GetDBInstance().Select("Token", "CreatedAt").Create(&arrobj).Error; err != nil {

		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}
