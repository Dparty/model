package restaurant

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/Dparty/common/utils"
	"gorm.io/gorm"
)

type Pair struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type Order struct {
	gorm.Model
	Item          Item   `json:"item" gorm:"type:JSON"`
	Specification []Pair `json:"specification"`
}

func (o Order) Equal(order Order) bool {
	if o.ID != order.ID {
		return false
	}
	om := o.SpecificationToMap()
	tm := order.SpecificationToMap()
	if len(om) != len(tm) {
		return false
	}
	for k, v := range om {
		if tm[k] != v {
			return false
		}
	}
	// for _, p := range o.
	return true
}

func (o Order) SpecificationToMap() map[string]string {
	var m map[string]string = make(map[string]string)
	for _, p := range o.Specification {
		m[p.Left] = p.Right
	}
	return m
}

func (o Order) Extra(p Pair) int64 {
	for _, attr := range o.Item.Attributes {
		if attr.Label == p.Left {
			for _, option := range attr.Options {
				if option.Label == p.Right {
					return option.Extra
				}
			}
		}
	}
	return 0
}

func (o Order) Total() int64 {
	var extra int64 = 0
	for _, option := range o.Specification {
		extra += o.Extra(option)
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
