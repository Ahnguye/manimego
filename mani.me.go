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


        buf := new(bytes.Buffer)
        buf.ReadFrom(r.Body)
        newStr := buf.String()

        fmt.Printf(newStr)
        //body, err := ioutil.ReadAll(r.Body)
        //writeError(err)

        //fmt.Printf("%s", body)
        insertDatabase(router.db, newStr)
        return
    } else if r.URL.Path == "/get" {
            queryDatabase(router.db, w, "SELECT * FROM orders;")
        return
    } 
    return
}

func queryDatabase(db *sql.DB, w http.ResponseWriter, query_str string) {
    rows, err := db.Query(query_str)
    writeError(err)

    for rows.Next() {
        var uid string
        var data string

        err = rows.Scan(&uid, &data)
        writeError(err)

        fmt.Fprintf(w, "%-10v %100v\n\n", uid, data)
    }

}

func insertDatabase(db *sql.DB, str string) {

    var lastInsertId int
    // QueryRow returns a row.
    //INSERT INTO orders (data) VALUES('{"name": "Paint house", "tags": ["Improvements", "Office"], "finished": false}');

    err := db.QueryRow("INSERT INTO orders(data) VALUES($1) returning uid;", str).Scan(&lastInsertId)
    writeError(err)
    fmt.Printf("%d\n", lastInsertId)
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

