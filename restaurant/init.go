package restaurant

import (
	"gorm.io/gorm"
)

var db *gorm.DB

var tableRepository TableRepository
var restaurantRepository RestaurantRepository

func Init(inject *gorm.DB) {
	db = inject
	tableRepository = NewTableRepository(db)
	restaurantRepository = NewRestaurantRepository(db)
	db.AutoMigrate(&Restaurant{}, &Item{}, &Printer{}, &Table{}, &Bill{})
}
