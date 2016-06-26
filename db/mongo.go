package db

import (
	"gopkg.in/mgo.v2"
)

func Connect(col string) (collection *mgo.Collection) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	return session.DB("gogame").C(col)
}
