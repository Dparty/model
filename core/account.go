package core

import (
	"github.com/Dparty/common/constants"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email    string         `json:"email" gorm:"index:email_index,unique"`
	Password string         `json:"password" gorm:"type:CHAR(128)"`
	Salt     []byte         `json:"salt"`
	Role     constants.Role `json:"role" gorm:"type:VARCHAR(128)"`
}
