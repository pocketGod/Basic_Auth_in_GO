package main

import (
	"net/http"

	"github.com/pocketGod/Basic_Auth_in_GO/Handlers"
)

func main() {
	http.HandleFunc("/tokenIssue", Handlers.TokenIssueHandler)
	http.HandleFunc("/tokenValidate", Handlers.TokenValidateHandler)
	http.ListenAndServe(":8080", nil)
}
