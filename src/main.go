package main

import (
	"fmt"
 //   "encoding/json"
    "log"
	"net/http"
	
    "github.com/gorilla/mux"
)
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
func createRecord(w http.ResponseWriter, r *http.Request) {}
func getRecord(w http.ResponseWriter, r *http.Request) {}

func main() {

    router := mux.NewRouter()
//	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/createRecord", createRecord).Methods("POST")
	router.HandleFunc("/getRecord/{id}",getRecord).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

