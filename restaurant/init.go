package restaurant

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(inject *gorm.DB) {
	db = inject
	db.AutoMigrate(&Restaurant{})
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&Printer{})
	db.AutoMigrate(&Table{})
	db.AutoMigrate(&Bill{})
}
