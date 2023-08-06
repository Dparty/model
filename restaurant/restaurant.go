package restaurant

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/Dparty/model/common"

	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	AccountId   uint
	Name        string
	Description string
	Items       []Item
}

type Item struct {
	gorm.Model
	RestaurantId uint
	Name         string
	Pricing      int64
	Properties   ItemProperties
}

type ItemProperties []ItemProperty

func (ItemProperties) GormDataType() string {
	return "JSON"
}

func (s *ItemProperties) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s ItemProperties) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type ItemProperty struct {
	Label  string
	Values common.StringList
}

func (ItemProperty) GormDataType() string {
	return "JSON"
}

func (s *ItemProperty) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s ItemProperty) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}
