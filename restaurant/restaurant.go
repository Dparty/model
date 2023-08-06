package restaurant

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	AccountId   uint
	Name        string
	Description string
	Items       []Item
}

type Item struct {
	gorm.Model
	Name       string
	Pricing    int64
	Properties []ItemProperty
}

type ItemProperty struct {
	Label  string
	Values []string
}

type Order struct {
	gorm.Model
	Items []Item
}
