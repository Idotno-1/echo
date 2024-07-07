package services

import "gorm.io/gorm"

var dbConn *gorm.DB

func SetDB(db *gorm.DB) {
	dbConn = db
}
