package main

import (
	"fmt"
	"net/http"
	"os/user"
)

// GetUser : get the current user
func GetUser() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Current User
	// fmt.Println("Username: " + user.Username)
	return user.Username

}
func main() {
	fmt.Println(GetUser())
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
	// GetUser()

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Username: %s", GetUser())
}
