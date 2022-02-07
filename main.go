package main

import (
	"github.com/gorilla/mux"
	"inputFileProject/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/upload", handler.FetchDataFromFile).Methods("POST")
	log.Fatal(http.ListenAndServe(":8099", router))
}
