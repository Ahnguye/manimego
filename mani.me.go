package main

import (
    "database/sql"
    "fmt"
    "net/http"
    _ "github.com/lib/pq"
    //"io/ioutil"
    "bytes"
)

const (  
  DB_HOST     = "manime.cyssztdd4zzm.us-west-2.rds.amazonaws.com"
  DB_PORT     = 5432
  DB_USER     = "ahnguye"
  DB_PASSWORD = "postgres"
  DB_NAME     = "manime"
)

type MyRouter struct {
    db *sql.DB
}



func (router *MyRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    r.ParseForm();
    if r.URL.Path == "/add" {
        fmt.Fprintf(w, "Sending message to client through customized router")

        return
    } else if r.URL.Path == "/get" {
	fmt.Fprintf(w, "Sending message to client through customized router")
         
        return
    } 
    return
}

func writeError(err error) {
    if err != nil {
        panic(err)
    
    }
}

func main() {
    fmt.Printf("backend in go\n")

    // Connect
    dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
                            DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    writeError(err)
    // Ping
    err = db.Ping()
    writeError(err)

    defer db.Close()

    err = http.ListenAndServe(":9092", &MyRouter{db})
    writeError(err)
    
}

/*

CREATE TABLE orders (
  uid SERIAL PRIMARY KEY,
  data text
);

*/

