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

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = utils.GenerteId()
	return err
}

func FindAccount(id uint) Account {
	var account Account
	db.Find(&account, id)
	return account
}

func FindAccountByEmail(email string) (account *Account) {
	db.Find(&account, "email = ?", email)
	return account
}
