package restaurant

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(inject *gorm.DB) {
	db = inject
	db.AutoMigrate(&Restaurant{}, &Item{}, &Printer{}, &Table{}, &Bill{})
}
