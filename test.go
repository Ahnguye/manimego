package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
	//"encoding/json"
	//"time"
	//"strconv"
)

const (  
  DB_HOST     = "messagedb.cyssztdd4zzm.us-west-2.rds.amazonaws.com"
  DB_PORT     = 5432
  DB_USER     = "ahnguye"
  DB_PASSWORD = "postgres"
  DB_NAME     = "messagedb"
)

type Message struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Time string `json:"time"`
    Message string `json:"message"`
    Latitude string `json:"latitude"`
    Longitude string `json:"longitude"`
    Views string `json:"views"`
}

type Messageslice struct {
    Messages []Message
}

type MyRouter struct {
	db *sql.DB
}

func (router *MyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();

	
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Sending message to client through customized router")
		return
	} else if r.URL.Path == "/add" {


	} 

	http.NotFound(w,r)
	return
}


func writeError(err error) {
	if err != nil {
		panic(err)
	
	}
	
}

func main() {
    fmt.Printf("hotspot backend in go\n")

    // Connect
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    						DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    writeError(err)
    // Ping
    err = db.Ping()
    writeError(err)

    defer db.Close()

    // Second argument of ListenAndServe takes a customized router
    // The customized router "MyRouter" needs to implement ( *) ServeHTTP(w, r). 
    // IOW, needs to implement Handler
    err = http.ListenAndServe(":9092", &MyRouter{db})
    writeError(err)
	
}

