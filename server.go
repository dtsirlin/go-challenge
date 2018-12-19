package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Display a single data
func ValidateUser(w http.ResponseWriter, r *http.Request) {

	// params := mux.Vars(r)
	// to do:
	// search the userId param to see if it matches the id of any users
	// if user not found, say invalid details (don't say user not found for better security)
	// if yes, see if the password is correct
	// if yes, spit out jwt
	// else say invalid details

	// for _, item := range people {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(&Person{})
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/validate/{userId}/{userPassword}", ValidateUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
