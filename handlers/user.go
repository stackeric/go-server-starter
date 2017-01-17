package user

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/stackeric/gostarter/db"
)

// List user
func List(c *gin.Context) {
	result := bson.M{}
	collect := "mycollection"
	err := db.Mdb.C(collect).Find(bson.M{}).One(&result)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"user": result,
	})
}
