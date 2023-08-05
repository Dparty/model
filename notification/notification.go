package notification

import (
	"gitea.svc.boardware.com/bwc/model/common"
	"gorm.io/gorm"
)

type NotificationType string

type Email struct {
	gorm.Model
	Sender          string            `json:"sender"`
	Receivers       common.StringList `json:"receivers"`
	CarbonCopy      common.StringList `json:"carbonCopy"`
	BlindCarbonCopy common.StringList `json:"blindCarbonCopy"`
	Message         string            `json:"message"`
}
