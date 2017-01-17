package db

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo

	// Mdb is db form server
	Mdb *mgo.Database
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	MongoDBUrl = "mongodb://username:password@localhost:27017/admin"
)

// Connect connects to mongodb
func Connect() {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	Session = s
	Mongo = mongo
	cs := Session.Clone()
	defer cs.Close()
	Mdb = Session.DB("mydb")
}
