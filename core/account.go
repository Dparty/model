package core

import (
	"github.com/Dparty/common/constants"
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

func (a Account) ID() uint {
	return a.Model.ID
}

func (a Account) Own(asset Asset) bool {
	return a.ID() == asset.Owner().ID()
}

type Asset interface {
	Owner() Account
	Delete() error
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = utils.GenerteId()
	return err
}
