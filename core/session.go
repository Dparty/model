package core

import (
	"time"

	"gitea.svc.boardware.com/bwc/common/constants"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	AccountId   uint                  `json:"accountId" gorm:"index:accountId_index"`
	Token       string                `json:"token"`
	TokeType    constants.TokenType   `json:"tokenType" gorm:"type:VARCHAR(128)"`
	TokenFormat constants.TokenFormat `json:"tokenFormat" gorm:"type:VARCHAR(128)"`
	ExpiredAt   time.Time             `json:"expiredAt"`
}
