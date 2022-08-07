package models

type Product struct {
	Name        string   `json:"name" bson:"name"`
	Type        string   `json:"type" bson:"type"`
	Price       int      `json:"price" bson:"price"`
	Description string   `json:"description" bson:"description"`
	Category    []string `json:"category" bson:"category"`
}
