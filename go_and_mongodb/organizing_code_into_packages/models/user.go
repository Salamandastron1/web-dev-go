package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Name   string        `json:"name"`
	Gender string        `json:"gender"`
	Age    int           `json:"age"`
	ID     bson.ObjectId `json:"id" bson:"_id"`
}
