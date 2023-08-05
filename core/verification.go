package core

import (
	"github.com/Dparty/common/constants"
	"gorm.io/gorm"
)

type VerificationCode struct {
	gorm.Model
	Identity string                            `json:"email" gorm:"index:verification_index"` // Email or phone number
	Purpose  constants.VerificationCodePurpose `json:"purpose" gorm:"type:VARCHAR(128)"`
	Code     string                            `json:"code" gorm:"type:CHAR(6)"`
	Tries    int64
}
