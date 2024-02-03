package main

import (
	"log"
	"net/http"
	"time"
)

var (
	// Version is the version of the server.
	Version = "0.0.1"
	ch1 = make(chan int)

	workers = 50
	errVar error
)

// workerFxn is a Go function that performs a specific task.
//
// It does not take any parameters and does not return any values.
func workerFxn() {
	for i:=range ch1{
		time.Sleep(time.Duration(i) * time.Second)
		log.Printf("Worker %d done", i)
	}
}

// handleHello is a Go function that handles the "Hello" endpoint.
//
// It takes in an http.ResponseWriter and an http.Request as parameters.
func handleHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the server"))

	go func() {
		ch1<-time.Now().Nanosecond()
	}()
}

// main is the entry point function for the program.
//
// It does not take any parameters and does not return anything.
func main() {
	for i:=0;i<workers;i++{
		go workerFxn()
	}
	http.HandleFunc("/hello", handleHello)
	log.Printf("Starting server on version %s", Version)
	log.Fatal(http.ListenAndServe(":8080", nil))

	if errVar!=nil{
		log.Fatalln("Error starting server: ", errVar)
	}
}