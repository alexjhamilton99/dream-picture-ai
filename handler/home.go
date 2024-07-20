package handler

import (
	"dream-picture-ai/view/home"
	// "fmt"
	"net/http"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	// return fmt.Errorf("Failed to generate picture")
	return home.Index().Render(r.Context(), w)
}
