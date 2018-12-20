package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

//User type contains credentials for a user
type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

var users []User

var hmacSampleSecret = []byte("danssecret")

//AddUsers adds a set of users to the users, an array of User
func AddUsers() {
	// add user credentials for new user with follow statements:
	// users = append(users, User{Username: <username>, Password: <password>})
	users = append(users, User{Username: "dantsirlin@gmail.com", Password: "testpass"})
	users = append(users, User{Username: "seconduser@gmail.com", Password: "password"})
}

//ValidateUser validates that input parameters of username and password are valid credentials and returns access and refresh token for credentials
func ValidateUser(w http.ResponseWriter, r *http.Request) {
	// if username parameter matches a username field for a user in list of users:
	// -- check if password parameter matches password for given user
	// -- if yes, write to responsewriter access token valid for 24hr and refresh token valid for 7 days
	// -- if no, write to responsewriter "Invalid credentials"
	// if username parameter matches no users, write to responsewriter "Invalid credentials"
	params := mux.Vars(r)
	foundUser := false
	for _, user := range users {
		if user.Username == params["username"] {
			// json.NewEncoder(w).Encode(item)
			if user.Password == params["password"] {
				fmt.Fprintf(w, "Access granted to "+params["username"]+"\n")
				accessTokenString, accessErr := NewAccessToken(params["username"])
				if accessErr != nil {
					fmt.Fprintf(w, accessErr.Error())
				}
				refreshTokenString, refreshErr := NewRefreshToken(params["username"])
				if refreshErr != nil {
					fmt.Fprintf(w, refreshErr.Error())
				}
				fmt.Fprintln(w, "access token:\n"+accessTokenString)
				fmt.Fprintln(w, "refresh token:\n"+refreshTokenString)
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

//NewAccessToken creates a new token with that expires in 24 hours from creation
func NewAccessToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      "access",
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Unix() + (60 * 60 * 24),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}

//NewRefreshToken creates a new token with that expires in 7 days from creation
func NewRefreshToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      "refresh",
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Unix() + (7 * 60 * 60 * 24),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	AddUsers()
	// ideally, there any other routes would display the instructions of how to use the two routes below
	router.HandleFunc("/validate/{username}/{password}", ValidateUser).Methods("GET")
	// todo: a route that takes in jwt token and loads screen showing welcome to that user/pass
	// router.HandleFunc("/authjwt", ValidateJWT).Methods.("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
