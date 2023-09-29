package restaurant

import (
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init(inject *gorm.DB) {
	db = inject
	fmt.Println("db:", db)
	db.AutoMigrate(&Restaurant{})
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&Printer{})
	db.AutoMigrate(&Table{})
	db.AutoMigrate(&Bill{})
}
