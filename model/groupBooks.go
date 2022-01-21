package model

type GroupBook struct {
	ID         string `gorm:"primaryKey";json:"id" bson:"_id,omitempty"`
	CategoryID string `json:"category_id,omitempty" bson:"category_id,omitempty"`
	BookID     string `json:"book_id,omitempty" bson:"book_id,omitempty"`
}
