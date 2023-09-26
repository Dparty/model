package restaurant

import (
	"github.com/Dparty/common/utils"
	"github.com/Dparty/model/core"
	"gorm.io/gorm"
)

type PrinterType string

const (
	BILL    PrinterType = "BILL"
	KITCHEN PrinterType = "KITCHEN"
)

type Printer struct {
	gorm.Model
	RestaurantId uint
	Name         string      `json:"name"`
	Sn           string      `json:"sn"`
	Description  string      `json:"description"`
	Type         PrinterType `json:"type" gorm:"type:VARCHAR(128)"`
}

func (p Printer) Owner() core.Account {
	return FindRestaurant(p.RestaurantId).Owner()
}

func (printer *Printer) BeforeCreate(tx *gorm.DB) (err error) {
	printer.ID = utils.GenerteId()
	return err
}
