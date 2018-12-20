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

var hmacSecret = []byte("danssecret")

//AddUsers adds a set of users to the users, an array of User
func AddUsers() {
	// add user credentials for new user with follow statements:
	// users = append(users, User{Username: <username>, Password: <password>})
	users = append(users, User{Username: "dantsirlin@gmail.com", Password: "testpass"})
	users = append(users, User{Username: "seconduser@gmail.com", Password: "password"})
	users = append(users, User{Username: "a-real-person@gmail.com", Password: "1234567890"})
}

//ValidateUser validates that input parameters of username and password are valid credentials and returns access and refresh token for credentials
func ValidateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	foundUser := false
	for _, user := range users {
		if user.Username == params["username"] {
			if user.Password == params["password"] {
				fmt.Fprintf(w, "Access granted to "+params["username"]+"\n")

				accessTokenString, accessErr := NewAccessToken(params["username"])
				fmt.Fprintf(w, "access token:\n")

				if accessErr != nil {
					fmt.Fprintf(w, accessErr.Error())
				} else {
					fmt.Fprintln(w, accessTokenString)
				}

				refreshTokenString, refreshErr := NewRefreshToken(params["username"])
				fmt.Fprintf(w, "refresh token:\n")

				if refreshErr != nil {
					fmt.Fprintf(w, refreshErr.Error())
				} else {
					fmt.Fprintln(w, refreshTokenString)
				}
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
		"iat":      float64(time.Now().Unix()),
		"exp":      float64(time.Now().Unix() + (60 * 60 * 24)),
	})

	tokenString, err := token.SignedString(hmacSecret)

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

	tokenString, err := token.SignedString(hmacSecret)

	return tokenString, err
}

//ValidateJWT check is jwt is valid if so, load protected endpoint
func ValidateJWT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tokenString := params["tokenString"]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	// error exists if not a valid jwt token
	// - incorrect format (3 strings separated by 2 periods and decodes correctly)
	// - expired token
	//
	// return "Access Denied" for any processing of jwt token that isn't a valid token giving access to valid user
	// - invalid token (reasons above)
	// - valid but not access token
	// - valid and access token, but not associated with a valid user
	if err != nil {
		fmt.Fprintf(w, "Access Denied")
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// big break through!  would've been nice if i knew that part of checking for valid token
			// is that if claim "exp" exists and that datetime is less than current time, thus expired token,
			// the token counts as invalid.
			//
			// my catcher for comparing the expiration time to current time only worked when i had the claim name
			// differ to "exp".  this catcher will be in a non-final commit but be ommitted later

			if claims["sub"] == "access" {
				userFound := false
				for _, user := range users {
					if user.Username == claims["username"] {
						fmt.Fprintf(w, "Access granted, "+claims["username"].(string)+"!")
						userFound = true
					}
				}

				if !userFound == true {
					fmt.Fprintf(w, "Access denied.")
				}
			} else {
				fmt.Fprintf(w, "Access denied.")
			}
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	AddUsers()
	// ideally, there any other routes would display the instructions of how to use the two routes below
	router.HandleFunc("/validate/{username}/{password}", ValidateUser).Methods("GET")
	router.HandleFunc("/authjwt/{tokenString}", ValidateJWT).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
