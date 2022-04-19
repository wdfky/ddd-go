package entity

import "github.com/google/uuid"

type Item struct {
	Id   uuid.UUID `json:"id" bson:"id"`
	Name string    `json:"name" bson:"name"`
	Desc string    `json:"desc" bson:"desc"`
}
