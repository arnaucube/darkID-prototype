package main

import (
	mgo "gopkg.in/mgo.v2"
)

func getSession() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://" + config.Mongodb.IP)
	if err != nil {
		panic(err)
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session, err
}
func getCollection(session *mgo.Session, collection string) *mgo.Collection {

	c := session.DB(config.Mongodb.Database).C(collection)
	return c
}
