package common

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type ObjectStorageProvider string

const (
	COS ObjectStorageProvider = "COS"
)

type ObjectStorage struct {
	gorm.Model
	Provider ObjectStorageProvider `json:"provider" gorm:"type:VARCHAR(128)"`
	Url      string                `json:"url"`
}

type List[T any] struct {
	Data []T
}

func (l List[T]) GormDataType() string {
	return "JSON"
}

func (l *List[T]) Scan(value any) error {
	return json.Unmarshal(value.([]byte), l.Data)
}

func (l List[T]) Value() (driver.Value, error) {
	b, err := json.Marshal(l.Data)
	return b, err
}

func (ObjectStorage) GormDataType() string {
	return "JSON"
}

func (s *ObjectStorage) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s ObjectStorage) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}
