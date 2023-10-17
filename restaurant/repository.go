package restaurant

import (
	abstract "github.com/Dparty/model/abstract"
	"gorm.io/gorm"
)

type RestaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return RestaurantRepository{
		db: db,
	}
}

func (r RestaurantRepository) Get(conds ...any) *Restaurant {
	var restaurant Restaurant
	ctx := r.db.Find(restaurant, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &restaurant
}

func (r RestaurantRepository) GetById(id uint) *Restaurant {
	return r.Get(id)
}

func (r RestaurantRepository) List(conds ...any) []Restaurant {
	var restaurants []Restaurant
	r.db.Find(&restaurants, conds...)
	return restaurants
}

func (r RestaurantRepository) ListBy(accountId *uint) []Restaurant {
	ctx := r.db.Model(&Restaurant{})
	if accountId != nil {
		ctx.Where("account_id = ?", accountId)
	}
	var restaurants []Restaurant
	ctx.Find(&restaurants)
	return restaurants
}

func (r RestaurantRepository) Create(owner abstract.Owner, name, description string) Restaurant {
	restaurant := Restaurant{
		Name:        name,
		Description: description,
	}
	restaurant.SetOwner(owner)
	db.Save(&restaurant)
	return restaurant
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return TableRepository{db}
}

type TableRepository struct {
	db *gorm.DB
}
