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
	Tables      []Table
}

type Item struct {
	gorm.Model
	RestaurantId uint
	Name         string
	Pricing      int64
	Attributes   Attributes
	Images       common.StringList `json:"images" gorm:"type:JSON"`
}

type Attributes []Attribute

func (Attributes) GormDataType() string {
	return "JSON"
}

func (s *Attributes) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s Attributes) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type Attribute struct {
	Label   string   `json:"label"`
	Options []Option `json:"options"`
}

type Option struct {
	Label string `json:"label"`
	Extra int64  `json:"extra"`
}

func (Attribute) GormDataType() string {
	return "JSON"
}

func (s *Attribute) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s Attribute) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type Table struct {
	gorm.Model
	RestaurantId uint
	Label        string `json:"label"`
}
