package core

import (
	"github.com/Dparty/common/constants"
	"github.com/Dparty/common/fault"
	"github.com/Dparty/common/utils"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email    string         `json:"email" gorm:"index:email_index,unique"`
	Password string         `json:"password" gorm:"type:CHAR(128)"`
	Salt     []byte         `json:"salt"`
	Role     constants.Role `json:"role" gorm:"type:VARCHAR(128)"`
}

func (a Account) Owner(asset Asset) bool {
	return a.ID == asset.Owner().ID
}

type Asset interface {
	Owner() Account
	Delete() error
}

func (a Account) DeleteAsset(asset Asset) error {
	if asset.Owner().ID != a.ID {
		return fault.ErrPermissionDenied
	}
	return asset.Delete()
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return err
}

func FindAccount(conds ...any) (account Account, err error) {
	ctx := db.Find(&account, conds...)
	if ctx.RowsAffected == 0 {
		return account, fault.ErrNotFound
	}
	return account, nil
}

func FindAccountByEmail(email string) (account Account, err error) {
	return FindAccount("email = ?", email)
}
