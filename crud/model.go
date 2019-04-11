package model

import (
	"errors"
	"fmt"
	"os"
	"time"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Required ENVs
// mdbConn 			- mongo connection string
// mdbName 			- mongo database name
// mdbCollection 	- mongo collection name

// CRUD model for using crud methods
// - required DB, DBname, DBCollection
type CRUD struct {
	ID   string
	Err  error
	Data interface{}
}

// mgoSession session for mongoDB
var mgoSession *mgo.Session

// MongoSession generate session for mongoDB
func MongoSession() *mgo.Session {

	if mgoSession == nil {

		var err error
		mgoSession, err = mgo.Dial(os.Getenv("mdbConn"))
		if err != nil {
			fmt.Println(err)
		}

		mgoSession.SetSocketTimeout(1 * time.Hour)
		mgoSession.SetMode(mgo.Monotonic, true)

	}
	return mgoSession.Clone()
}

// Create insert model.Data
func (model *CRUD) Create() {

	// Generate mongoDB session
	DB := MongoSession()
	defer DB.Close() // Close when done

	// Generate new ID
	model.ID = bson.NewObjectId().Hex()

	_, err := DB.
		DB(os.Getenv("mdbConn")).
		C(os.Getenv("mdbCollection")).
		UpsertId(bson.ObjectIdHex(model.ID), model.Data)
	if err != nil {
		model.Err = err
	}

	return
}

// Read return Data from ID
// - required ID
func (model *CRUD) Read() {

	// Generate mongoDB session
	DB := MongoSession()
	defer DB.Close() // Close when done

	// Check if ID not empty
	if model.ID == "" {
		model.Err = errors.New("ID empty")
		return
	}

	err := DB.
		DB(os.Getenv("mdbConn")).
		C(os.Getenv("mdbCollection")).
		Find(bson.M{"_id": bson.ObjectIdHex(model.ID)}).
		One(&model.Data)

	if err != nil {
		model.Err = err
		if err.Error() == "not found" {
			model.Data = nil
		}

	}

	return
}

// Update update by ID, data in Data
// - required ID
func (model *CRUD) Update() {

	// Generate mongoDB session
	DB := MongoSession()
	defer DB.Close() // Close when done

	if model.ID == "" {
		model.Err = errors.New("ID empty")
		return
	}

	err := DB.
		DB(os.Getenv("mdbConn")).
		C(os.Getenv("mdbCollection")).
		Update(
			bson.M{"_id": bson.ObjectIdHex(model.ID)},
			bson.M{"$set": model.Data},
		)
	if err != nil {
		model.Err = err
	}

}

// Delete row in db by ID
// - required ID
func (model *CRUD) Delete() {

	// Generate mongoDB session
	DB := MongoSession()
	defer DB.Close() // Close when done

	if model.ID == "" {
		model.Err = errors.New("ID empty")
		return
	}

	err := DB.
		DB(os.Getenv("mdbConn")).
		C(os.Getenv("mdbCollection")).
		RemoveId(bson.ObjectIdHex(model.ID))
	if err != nil {
		model.Err = err
	}
	return
}
