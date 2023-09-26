package model

import (
	"fmt"

	"github.com/Dparty/model/core"
	"github.com/Dparty/model/restaurant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(user, password, host, port, database string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func Init(db *gorm.DB) {
	core.Init(db)
	restaurant.Init(db)
}

type Asset interface {
	Owner() core.Account
}
