package restaurant

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/Dparty/common/utils"
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
	Printers    []Printer
	Categories  []Category
}

type Category struct {
	gorm.Model
	RestaurantId uint
	Name         string
}

func (r *Restaurant) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = utils.GenerteId()
	return
}

type Item struct {
	gorm.Model
	RestaurantId uint
	Name         string
	Pricing      int64
	Attributes   Attributes
	Images       common.StringList `json:"images" gorm:"type:JSON"`
	Tags         common.StringList
	Printers     common.IDList
	Categories   common.IDList
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = utils.GenerteId()
	return
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

type Options []Option

func (Options) GormDataType() string {
	return "JSON"
}

func (s *Options) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s Options) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
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

func (t *Table) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = utils.GenerteId()
	return
}

type OrderItem struct {
	gorm.Model
	Name    string  `json:"itemsName"`
	Pricing int64   `json:"pricing"`
	Options Options `json:"options"`
}

func (o OrderItem) Total() int64 {
	var extra int64 = 0
	for _, option := range o.Options {
		extra += option.Extra
	}
	return o.Pricing + extra
}

type OrderItems []OrderItem

func (OrderItems) GormDataType() string {
	return "JSON"
}

func (s *OrderItems) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s OrderItems) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type Bill struct {
	gorm.Model
	Items       OrderItems
	TableLabel  string
	CheckoutUrl string
}

func (b Bill) Total() int64 {
	var total int64 = 0
	for _, item := range b.Items {
		total += item.Total()
	}
	return total
}

func (b *Bill) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = utils.GenerteId()
	return err
}

type Printer struct {
	gorm.Model
	RestaurantId uint
	Name         string `json:"name"`
	Sn           string `json:"sn"`
	Description  string `json:"description"`
}

func (printer *Printer) BeforeCreate(tx *gorm.DB) (err error) {
	printer.ID = utils.GenerteId()
	return err
}
