package restaurant

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

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
	db.Model(&Restaurant{}).Preload("Items").Preload("Tables").Preload("Printers").Find(&r, id)
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
	Tags        common.StringList
}

func (r Restaurant) GetItems() []Item {
	var items []Item
	db.Where("restaurant_id = ?", r.ID).Find(&items)
	return items
}

func (r Restaurant) ListBill(startAt, endAt *time.Time) []Bill {
	var bills []Bill
	ctx := db.Model(&bills)
	if startAt != nil {
		ctx = ctx.Where("created_at >= ?", startAt)
	}
	if endAt != nil {
		ctx = ctx.Where("created_at <= ?", endAt)
	}
	ctx.Find(&bills)
	return bills
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

type Attributes []Attribute

func (as Attributes) GetOption(left, right string) (Pair, error) {
	for _, a := range as {
		if left == a.Label {
			for _, option := range a.Options {
				if right == option.Label {
					return Pair{Left: left, Right: right}, nil
				}
			}
		}
	}
	return Pair{}, errors.New("NotFound")
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
