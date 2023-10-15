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

func (r Repository) Find(account *Account, conds ...any) (*Account, *gorm.DB) {
	ctx := r.db.Find(account, conds...)
	if ctx.RowsAffected == 0 {
		return nil, ctx
	}
	return account, ctx
}

func Find(conds ...any) (Account, error) {
	account, err := model.Find(&Account{}, conds...)
	return *account, err
}

func FindAccountByEmail(email string) (account Account, err error) {
	return Find("email = ?", email)
}
