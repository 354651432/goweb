package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"userName"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	LoginTime string `json:"loginTime"`
	LoginIp   string `json:"loginIp"`
	Skin      string `json:"skin"`
}

func (_ *User) TableName() string {
	return "hz_user"
}

var mysqlDns = ""

func SetDsn(dns string) {
	mysqlDns = dns
}

func Open() *gorm.DB {
	var err error
	var db *gorm.DB
	db, err = gorm.Open("mysql", mysqlDns)
	if err != nil {
		log.Println(err)
	}

	return db
}
