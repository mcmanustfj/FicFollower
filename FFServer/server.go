// A basic web server in Go

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path[1:])
}

func resetdb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Resetting DB")
	//if foo.db exists, delete it
	if _, err := os.Stat("./ff.db"); err == nil {
		os.Remove("./ff.db")
	}
	db, err := sql.Open("sqlite3", "./ff.db")
	if err != nil {
		fmt.Fprintf(w, "Error, see logs")
		fmt.Println(err)
		return
	}
	sqlStmt := "select 1;"
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Fprintf(w, "Error, see logs")
		fmt.Println(err)
		return
	}

}

func main() {
	http.HandleFunc("/resetdb", resetdb)

	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	http.HandleFunc("/", home)

	// ListenAndServe listens on the TCP network address addr and
	// then calls Serve with handler to handle requests on incoming connections.
	http.ListenAndServe(":8080", nil)
}
