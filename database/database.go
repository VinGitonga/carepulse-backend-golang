package database

import "gorm.io/gorm"

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance
