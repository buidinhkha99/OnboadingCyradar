package model

type Book struct {
	ID          string  `json:"id" bson:"_id,omitempty"`
	Name        string  `json:"name,omitempty" bson:"name,omitempty"`
	Quantily    int     `json:"quantily,omitempty" bson:"quantily,omitempty"`
	Description string  `json:"description,omitempty" bson:"description,omitempty"`
	Price       float32 `json:"price,omitempty" bson:"price,omitempty"`
	Rate        float32 `json:"rate,omitempty" bson:"rate,omitempty"`
	Image       string  `json:"image,omitempty" bson:"image,omitempty"`
}

type DetailBook struct {
	Book       Book        `json:"book" bson:"book,omitempty"`
	Category   []Category  `json:"category" bson:"category,omitempty"`
	GroupBooks []GroupBook `json:"group" bson:"group,omitempty"`
}
