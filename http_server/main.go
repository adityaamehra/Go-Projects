/*
200's :- The code is perfect, and it is working
400's :- There is an error
500's :- Problem with the server

There are 4 major methods:-

1.) GET :- To get a response
2.) POST :- To put some data
3.) PUT :- To update the data
4.) DELETE :- To delete the data
*/
package main

import (
	"fmt"
	"net/http"
)

func eventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.WriteHeader(http.StatusOK)
	if r.Method == "GET" {
		w.Write([]byte("Welcome to GoFR's event!"))
	}
	if r.Method == "POST" {
		w.Write([]byte("This is a POST method and it is good"))
	}
}

func main() {
	http.HandleFunc("/event", eventHandler)
	fmt.Println("Listening on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Server encounters an error : #{err}\n")
	}
}
