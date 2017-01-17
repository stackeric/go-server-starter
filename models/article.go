package models

import "gopkg.in/mgo.v2/bson"

const (
	// Collection holds the name of the articles collection
	Collection = "mycollection"
)

// Article model
type Article struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string        `json:"title" form:"title" binding:"required" bson:"title"`
	Body      string        `json:"body" form:"body" binding:"required" bson:"body"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`
	// User      bson.ObjectId `json:"user"`
}
