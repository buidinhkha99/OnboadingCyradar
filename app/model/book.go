package model

type Book struct {
	ID          uint64  `json:"id,omitempty" `
	Name        string  `json:"name,omitempty" `
	Quantily    int     `json:"quantily,omitempty" `
	Description string  `json:"description,omitempty" `
	Price       float32 `json:"price,omitempty" `
	Rate        float32 `json:"rate,omitempty" `
	Image       string  `json:"image,omitempty" `
}

type DetailBook struct {
	Book       Book        `json:"book"`
	Category   []Category  `json:"category"`
	GroupBooks []GroupBook `json:"group"`
}
