package model

type Product struct {
	ID          string  `json:"id" bson:"_id,omitempty"` // MongoDB uses _id
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Category    string  `json:"category" bson:"category"`
	Price       float64 `json:"price" bson:"price"`
	Brand       string  `json:"brand" bson:"brand"`
	Quantity    int     `json:"quantity" bson:"quantity"`
}
