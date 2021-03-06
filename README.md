# MicroAPI
CRUD API using GO&MongoDB

## Dependencies

* github.com/globalsign/mgo     - mongoDB driver
* github.com/gorilla/mux        - router
* github.com/gorilla/handlers   - handlers

## ENV Requirement

* mdbConn           - mongo connection string
* mdbName           - mongo database name
* mdbCollection     - mongo collection name
* apiPort           - api port

## Desc

API handle CRUD functions

```
muxRouter.Handle("/", CreateHandler).Methods("POST")
muxRouter.Handle("/{id}", ReadHandler).Methods("GET")
muxRouter.Handle("/{id}", UpdateHandler).Methods("PUT")
muxRouter.Handle("/{id}", DeleteHandler).Methods("DELETE")
```

For Create&Update send data in form-value 
- key:body
- value:JSON object

Sample model for mongoDB in main.js
- change to your model
```
type ExampleData struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name,omitempty"`
	Active string `json:"active" bson:"active,omitempty"`
}
```

## Run

```
mdbConn="localhost:27017" mdbName="foo" mdbCollection="bar" apiPort="80" go run main.go
```

## Docker

```
docker build -t microapi .
```
```
docker run -d --name microapi -p 9900:80 -e "mdbConn=localhost:27017" -e "mdbName=foo" -e "mdbCollection=bar" -e "apiPort=80" microapi
```

