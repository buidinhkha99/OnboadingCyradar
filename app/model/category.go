package model

type Category struct {
	ID   uint64 `gorm:"primaryKey";json:"id, omitempty"`
	Name string `json:"name, omitempty"`
}
