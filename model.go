package model

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var ErrNotFound = errors.New("entity not found")

func NewConnection(user, password, host, port, database string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func Find[T any](model T, conds ...any) (T, error) {
	if ctx := db.Find(model, conds...); ctx.RowsAffected == 0 {
		return model, nil
	}
	return model, nil
}

func Init(inject *gorm.DB) {
	db = inject
}
