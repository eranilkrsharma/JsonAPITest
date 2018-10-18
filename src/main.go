package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"math/rand"
    "github.com/gorilla/mux"
)

type Record struct {
	recordtype string
	recordCreator string
	id int
	 
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
func createRecord(w http.ResponseWriter, r *http.Request) {
	record := Record{} //create empty record
	fmt.Fprintf(w, json.NewDecoder(r.Body).Decode(&record).toString())
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil{
		panic(err)
	}

	record.id = rand.Int()
    fmt.Fprintf(w, record.recordtype)
	responseJson, err := json.Marshal(record)
	fmt.Fprintf(w, "parsing done")
	if err != nil{
		panic(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)

}
func getRecord(w http.ResponseWriter, r *http.Request) {}

func main() {

    router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/createRecord", createRecord)
	router.HandleFunc("/getRecord/{id}",getRecord)
    log.Fatal(http.ListenAndServe(":8000", router))
}

