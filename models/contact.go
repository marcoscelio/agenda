package models

import "gopkg.in/mgo.v2/bson"

type Contact struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	Email  string        `bson:"email" json:"email"`
	Active bool          `bson:"active" json:"active"`
}
