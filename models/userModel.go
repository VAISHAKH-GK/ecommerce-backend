package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name"  bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"-"     bson:"password"`
}

type AdminUser struct {
	UserName string `json:"userName" bson:"userName"`
	Password string `json:"password"  bson:"password"`
}
