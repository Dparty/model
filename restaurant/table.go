package restaurant

import (
	"github.com/Dparty/common/fault"
	"github.com/Dparty/common/utils"
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	RestaurantId uint
	Label        string `json:"label"`
	X            int64  `json:"x"`
	Y            int64  `json:"y"`
}

func (t Table) ListBills() []Bill {
	var bills []Bill
	db.Find(&bills, "table_id = ?", t.ID)
	return bills
}

func (t Table) Owner() Restaurant {
	return *restaurantRepository.GetById(t.RestaurantId)
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
