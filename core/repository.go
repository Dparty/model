package core

import (
	"github.com/Dparty/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Get(conds ...any) *Account {
	var account Account
	ctx := r.db.Find(account, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &account
}

func (r Repository) GetById(id uint) *Account {
	return r.Get(id)
}

func (r Repository) GetByEmail(email string) *Account {
	return r.Get("email = ?", email)
}

func Find(conds ...any) (Account, error) {
	account, err := model.Find(&Account{}, conds...)
	return *account, err
}

func FindAccountByEmail(email string) (account Account, err error) {
	return Find("email = ?", email)
}
