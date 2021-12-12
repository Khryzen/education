package api

import (
	"fmt"
	"net/http"
	"strings"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	api := strings.TrimSuffix(r.URL.Path, "/")

	fmt.Println("api", api)

	if strings.HasPrefix(api, "validate_user") {
		ValidateUserHandler(w, r)

	}

	if strings.HasPrefix(api, "validate_username") {
		fmt.Println("ValidateUsername START")
		ValidateUserNameHandler(w, r)
		fmt.Println("ValidateUsername END")

	}
}
