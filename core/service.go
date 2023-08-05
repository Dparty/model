package core

import (
	"github.com/Dparty/common/constants"

	"github.com/Dparty/common/utils"

	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name        string                `json:"name" gorm:"index:name_index,unique"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Url         string                `json:"url"`
	Type        constants.ServiceType `json:"type" gorm:"type:VARCHAR(128)"`
}

func NewService(name, title, description, url string, serviceType constants.ServiceType) Service {
	service := Service{
		Name:        name,
		Title:       title,
		Description: description,
		Url:         url,
		Type:        serviceType,
	}
	service.ID = utils.GenerteId()
	return service
}
