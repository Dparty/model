package restaurant

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/Dparty/common/utils"
	"github.com/Dparty/model/common"
	"github.com/Dparty/model/core"
	"gorm.io/gorm"
)

func FindItem(id any) *Item {
	var r *Item
	db.Model(&Item{}).Find(&r, id)
	return r
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

func (i Item) Extra(p Pair) int64 {
	for _, attr := range i.Attributes {
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
