package gorm

import (
	"database/sql"
	"fmt"

	"github.com/putranurf/be-tpd/db"
	"github.com/putranurf/be-tpd/helpers"
)

func CheckLogin(username, password string) (bool, error) {
	// var obj migration.User
	// var arrobj []migration.User
	// var res Response
	var pwd string

	if err := db.GetDBInstance().Raw("select * from master_user where username = ?", username).Error; err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Username not Found")
			return false, err
		}
		match, err := helpers.CheckPasswordHash(password, pwd)
		if !match {
			fmt.Println("Hash and Pass doesn;t match")
			return false, err
		}
		fmt.Println("Query Error")
		return false, err
	}

	// arrobj = append(arrobj, obj)

	// res.Status = http.StatusOK
	// res.Message = "Success"
	// res.Data = arrobj

	return true, nil

}
