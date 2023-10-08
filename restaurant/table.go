package restaurant

import (
	"github.com/Dparty/common/fault"
	"github.com/Dparty/common/utils"
	"github.com/Dparty/model/core"
	"gorm.io/gorm"
)

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

func FindTable(conds ...interface{}) (Table, error) {
	var table Table
	ctx := db.Find(&table, conds)
	if ctx.RowsAffected == 0 {
		return table, fault.ErrNotFound
	}
	return table, nil
}

func ListTable(conds ...interface{}) []Table {
	var tables []Table
	db.Find(&tables, conds)
	return tables
}
