package restaurant

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/Dparty/common/utils"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Item    Item    `json:"item" gorm:"type:JSON"`
	Options Options `json:"options"`
}

func (o Order) Total() int64 {
	var extra int64 = 0
	for _, option := range o.Options {
		extra += option.Extra
	}
	return o.Item.Pricing + extra
}

type Orders []Order

func (Orders) GormDataType() string {
	return "JSON"
}

func (s *Orders) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s Orders) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type Bill struct {
	gorm.Model
	RestaurantId uint `gorm:"index:rest_id"`
	Orders       Orders
	PickUpCode   int64
	TableLabel   string
}

func (b Bill) Total() int64 {
	var total int64 = 0
	for _, item := range b.Orders {
		total += item.Total()
	}
	return total
}

func (b *Bill) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = utils.GenerteId()
	return err
}
