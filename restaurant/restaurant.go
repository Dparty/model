package restaurant

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/Dparty/common/utils"
	"github.com/Dparty/model/common"
	"github.com/Dparty/model/core"
	"gorm.io/gorm"
)

func CreateRestaurant(accountId uint, name, description string) Restaurant {
	r := Restaurant{
		AccountId:   accountId,
		Name:        name,
		Description: description,
	}
	db.Save(&r)
	return r
}

func FindRestaurant(id uint) *Restaurant {
	var r *Restaurant
	db.Model(&Restaurant{}).Preload("Items").Preload("Tables").Preload("Printers").Find(&r)
	return r
}

type Restaurant struct {
	gorm.Model
	AccountId   uint
	Name        string
	Description string
	Items       []Item
	Tables      []Table
	Printers    []Printer
}

func (r Restaurant) Owner() core.Account {
	var account core.Account
	db.Find(&account, r.AccountId)
	return account
}

func (r *Restaurant) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = utils.GenerteId()
	return
}

func (r *Restaurant) AddPrinter(sn string, name string, t PrinterType) Printer {
	printer := Printer{
		RestaurantId: r.ID,
		Name:         name,
		Sn:           sn,
		Type:         t,
	}
	db.Save(&printer)
	r.Printers = append(r.Printers, printer)
	return printer
}

func (r *Restaurant) AddTable(label string) Table {
	table := Table{
		RestaurantId: r.ID,
		Label:        label,
	}
	db.Save(&table)
	r.Tables = append(r.Tables, table)
	return table
}

func (r *Restaurant) AddItem(item Item) Item {
	item.ID = r.ID
	db.Save(&item)
	r.Items = append(r.Items, item)
	return item
}

type Item struct {
	gorm.Model
	RestaurantId uint              `json:"restaurantId"`
	Name         string            `json:"name"`
	Pricing      int64             `json:"pricing"`
	Attributes   Attributes        `json:"attributes"`
	Images       common.StringList `json:"images" gorm:"type:JSON"`
	Tags         common.StringList `json:"tags"`
	Printers     common.IDList     `json:"printers"`
	Categories   common.IDList     `json:"categories"`
}

func (i Item) Owner() core.Account {
	return FindRestaurant(i.RestaurantId).Owner()
}

func (i *Item) AddAttribute(att Attribute) Item {
	i.Attributes = append(i.Attributes, att)
	return *i
}

func (s *Item) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s Item) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = utils.GenerteId()
	return
}

type Attributes []Attribute

func (as Attributes) GetOption(left, right string) (Option, error) {
	for _, a := range as {
		if left == a.Label {
			for _, option := range a.Options {
				if right == option.Label {
					return option, nil
				}
			}
		}
	}
	return Option{}, errors.New("NotFound")
}

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

func (t Table) Owner() core.Account {
	return FindRestaurant(t.RestaurantId).Owner()
}

func (t *Table) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = utils.GenerteId()
	return
}

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
	CheckoutUrl  string
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
