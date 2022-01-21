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
	Book       Book        `json:"book" `
	Category   []Category  `json:"category"`
	GroupBooks []GroupBook `json:"group" `
	Status     bool        `json:"status"`
	Books      []Book      `json:"books"`
}

type BookPublish struct {
	IdBook      string    `json:"id"`
	Description string    `json:"description"`
	Channel     string    `json:"channel"`
	GroupBook   GroupBook `json:"groupbook"`
}

type CatPublish struct {
	IdCatergory string    `json:"id"`
	Description string    `json:"description"`
	Channel     string    `json:"channel"`
	Book        []Book    `json:"book"`
	GroupBook   GroupBook `json:"groupbook"`
}
