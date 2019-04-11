package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	model "github.com/FilipAnteKovacic/microAPI/crud"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ENVs
// mdbConn 			- mongo connection string
// mdbName 			- mongo database name
// mdbCollection 	- mongo collection name
// apiPort			- api port

// ErrorResponse print if error given
type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

//SuccessResponse print if request ok
type SuccessResponse struct {
	Status string      `json:"status"`
	ID     string      `json:"id,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// ExampleData model for using crud methods
type ExampleData struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name,omitempty"`
	Active string `json:"active" bson:"active,omitempty"`
}

// ParseRequestValues parse request body to stuct
func ParseRequestValues(r *http.Request) (ExampleData, error) {

	body := ExampleData{}

	err := json.Unmarshal([]byte(r.FormValue("body")), &body)

	return body, err
}

// CreateHandler handle create requests on api
var CreateHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	InsertData, err := ParseRequestValues(r)
	// Check error
	if err != nil {
		Print(r, w, ErrorResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	// Setup CRUD model for insert
	InsertExample := model.CRUD{
		Data: InsertData,
	}

	// Create new row in DB
	InsertExample.Create()

	// Check error
	if InsertExample.Err != nil {
		Print(r, w, ErrorResponse{
			Status: "error",
			Error:  InsertExample.Err.Error(),
		})
		return
	}

	// Print ID on success
	Print(r, w, SuccessResponse{
		Status: "success",
		ID:     InsertExample.ID,
	})
	return

})

// ReadHandler handle all requests on api
var ReadHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// URL vars
	vars := mux.Vars(r)

	// Setup CRUD model for read
	ReadExample := model.CRUD{
		ID: vars["id"],
	}

	// Read from DB, using created ID
	ReadExample.Read()

	// Check error
	if ReadExample.Err != nil {
		Print(r, w, ErrorResponse{
			Status: "error",
			Error:  ReadExample.Err.Error(),
		})
		return
	}

	// Print ID&Data on success
	Print(r, w, SuccessResponse{
		Status: "success",
		Data:   ReadExample.Data,
	})
	return

})

// UpdateHandler handle all requests on api
var UpdateHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// URL vars
	vars := mux.Vars(r)

	UpdateData, err := ParseRequestValues(r)
	// Check error
	if err != nil {
		Print(r, w, ErrorResponse{Error: err.Error()})
		return
	}

	// Setup CRUD model for update
	UpdateExample := model.CRUD{
		ID:   vars["id"],
		Data: UpdateData,
	}

	// Update row in DB, using created ID
	UpdateExample.Update()

	// Check error
	if UpdateExample.Err != nil {
		Print(r, w, ErrorResponse{
			Status: "error",
			Error:  UpdateExample.Err.Error(),
		})
		return
	}

	// Print ID on success
	Print(r, w, SuccessResponse{
		Status: "success",
		ID:     UpdateExample.ID,
	})
	return

})

// DeleteHandler handle all requests on api
var DeleteHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// URL vars
	vars := mux.Vars(r)

	// Delete row in DB, using created ID
	DeleteExample := model.CRUD{
		ID: vars["id"],
	}

	DeleteExample.Delete()

	// Check error
	if DeleteExample.Err != nil {
		Print(r, w, ErrorResponse{
			Status: "error",
			Error:  DeleteExample.Err.Error(),
		})
		return
	}

	// Print ID on success
	Print(r, w, SuccessResponse{
		Status: "success",
		ID:     DeleteExample.ID,
	})
	return

})

// Print prints json response
func Print(r *http.Request, w http.ResponseWriter, res interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	jsonRsp, err := json.Marshal(res)

	if err != nil {

		jsonRsp, _ := json.Marshal(ErrorResponse{Error: "system error"})
		io.WriteString(w, string(jsonRsp))
		return

	}

	io.WriteString(w, string(jsonRsp))
	return
}

// StartAPI start server on port - ENV - apiPort
func main() {

	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.Handle("/", CreateHandler).Methods("POST")
	muxRouter.Handle("/{id}", ReadHandler).Methods("GET")
	muxRouter.Handle("/{id}", UpdateHandler).Methods("PUT")
	muxRouter.Handle("/{id}", DeleteHandler).Methods("DELETE")

	err := http.ListenAndServe(":"+os.Getenv("apiPort"), handlers.CompressHandler(muxRouter))
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}

}
