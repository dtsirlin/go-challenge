package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// The user Type
type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

var users []User

// Display a single data
func ValidateUser(w http.ResponseWriter, r *http.Request) {

	// to do:
	// search the userId param to see if it matches the id of any users
	// if user not found, say invalid details (don't say user not found for better security)
	// if yes, see if the password is correct
	// if yes, spit out jwt
	// else say invalid details

	params := mux.Vars(r)
	foundUser := false
	for _, user := range users {
		if user.Username == params["userUsername"] {
			// json.NewEncoder(w).Encode(item)
			if user.Password == params["userPassword"] {
				fmt.Fprintf(w, "Access granted to "+params["userUsername"]+":"+params["userPassword"])
				// todo: generate and print jwt token to screen
			} else {
				fmt.Fprintf(w, "Invalid credentials")
			}
			return
		}
	}
	if foundUser == false {
		fmt.Fprintf(w, "Invalid credentials")
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	users = append(users, User{Username: "dantsirlin@gmail.com", Password: "testpass"})
	users = append(users, User{Username: "seconduser@gmail.com", Password: "password"})
	// todo: route all pages that aren't /validate/username/pass to root page which shows user how to correctly use the rest api
	router.HandleFunc("/", Index)
	router.HandleFunc("/validate/{userUsername}/{userPassword}", ValidateUser).Methods("GET")
	// todo: a route that takes in jwt token and loads screen showing welcome to that user/pass
	// router.HandleFunc("/authjwt", ValidateJWT).Methods.("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
