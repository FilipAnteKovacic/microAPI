package model

import (
	"fmt"
)

// ENVs
// mdbConn 			- mongo connection string
// mdbName 			- mongo database name
// mdbCollection 	- mongo collection name

// MongoDBFooBar struct for testing model
type MongoDBFooBar struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name,omitempty"`
	Active string `json:"active" bson:"active,omitempty"`
}

func mongo() {

	// Init CRUD model, set row data
	InsertExample := CRUD{
		Data: MongoDBFooBar{
			Name:   "FooBar",
			Active: "true",
		},
	}

	// Create new row in DB
	fmt.Println("- Create")
	InsertExample.Create()

	if InsertExample.Err != nil {
		panic(InsertExample.Err)
	}

	NewID := InsertExample.ID
	fmt.Println(NewID)

	// Read from DB, using created ID
	ReadExample := CRUD{
		ID: NewID,
	}

	fmt.Println("- Read")
	ReadExample.Read()
	if ReadExample.Err != nil {
		panic(ReadExample.Err)
	}

	fmt.Println(ReadExample.Data)

	// Update row in DB, using created ID
	UpdateExample := CRUD{
		ID: NewID,
		Data: MongoDBFooBar{
			Name:   "FooBar",
			Active: "false",
		},
	}

	fmt.Println("- Update")
	UpdateExample.Update()
	if UpdateExample.Err != nil {
		panic(UpdateExample.Err)
	}

	// Read from DB, check if update worked
	fmt.Println("- Read")
	ReadExample.Read()
	if ReadExample.Err != nil {
		panic(ReadExample.Err)
	}
	fmt.Println(ReadExample.Data)

	// Delete row in DB, using created ID
	DeleteExample := CRUD{
		ID: NewID,
	}

	fmt.Println("- Delete")
	DeleteExample.Delete()
	if DeleteExample.Err != nil {
		panic(DeleteExample.Err)
	}

	// Read from DB, check if delete worked
	fmt.Println("- Read")
	ReadExample.Read()
	if ReadExample.Err != nil {
		panic(ReadExample.Err)
	}
	fmt.Println(ReadExample.Data)

}
