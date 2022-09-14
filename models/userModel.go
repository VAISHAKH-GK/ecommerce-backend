package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name"  bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"-"     bson:"password"`
}

type AdminUser struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserName   string             `json:"userName" bson:"userName"`
	Password   string             `json:"-"  bson:"password"`
	SuperAdmin bool               `json:"superAdmin" bson:"superAdmin"`
}

type UserCart struct {
	Id       primitive.ObjectID       `json:"_id" bson:"_id"`
	Products []map[string]interface{} `json:"products" bson:"products"`
	UserId   primitive.ObjectID       `json:"userId" bson:"userId"`
}
