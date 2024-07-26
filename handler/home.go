package handler

import (
	"dream-picture-ai/view/home"
	"fmt"
	"net/http"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)
	fmt.Printf("%+v\n", user.Account) // +v includes the field name and values
	return home.Index().Render(r.Context(), w)
}
