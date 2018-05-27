package config

import (
	"fmt"

	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

// DB Database
var DB *mgo.Database

// Books Collections
var Books *mgo.Collection

func init() {
	// Get Mongo Session
	// session, err := mgo.Dial("mongodb://username:password@localhost/bookstore")
	session, err := mgo.Dial("mongodb://localhost/bookstore")
	if err != nil {
		panic(err)
	}

	DB = session.DB("bookstore")
	Books = DB.C("books")

	fmt.Println("Connected to the Mongo Database Successfully ...")
}
