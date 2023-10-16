package restaurant

import (
	interfaces "github.com/Dparty/model/abstract"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// Owner implements interfaces.Asset.
func (Repository) Owner() interfaces.Owner {
	panic("unimplemented")
}

// Save implements interfaces.Asset.
func (Repository) Save() error {
	panic("unimplemented")
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Get(conds ...any) *Restaurant {
	var restaurant Restaurant
	ctx := r.db.Find(restaurant, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &restaurant
}

func (r Repository) GetById(id uint) *Restaurant {
	return r.Get(id)
}

func (r Repository) List(conds ...any) []Restaurant {
	var restaurants []Restaurant
	r.db.Find(&restaurants, conds...)
	return restaurants
}

func (r Repository) ListBy(accountId *uint) []Restaurant {
	ctx := r.db.Model(&Restaurant{})
	if accountId != nil {
		ctx.Where("account_id = ?", accountId)
	}
	var restaurants []Restaurant
	ctx.Find(&restaurants)
	return restaurants
}

func (r Repository) Create(owner interfaces.Owner, name, description string) Restaurant {
	restaurant := Restaurant{
		Name:        name,
		Description: description,
	}
	restaurant.SetOwner(owner)
	db.Save(&restaurant)
	return restaurant
}
