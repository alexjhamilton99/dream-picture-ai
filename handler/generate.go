package handler

import (
	"dream-picture-ai/db"
	"dream-picture-ai/types"
	"dream-picture-ai/view/generate"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleGenerateIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)
	images, err := db.GetImagesByUserID(user.ID)
	if err != nil {
		return err
	}
	data := generate.ViewData{
		Images: images,
	}
	return render(r, w, generate.Index(data))
}

func HandleGenerateCreate(w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)
	prompt := "Red sportscar in a garden"
	img := types.Image{
		Prompt: prompt,
		UserID: user.ID,
		Status: types.ImageStatusPending,
	}
	if err := db.CreateImage(&img); err != nil {
		return err
	}
	return render(r, w, generate.GalleryImage(img))
}

func HandleGenerateImageStatus(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}
	image, err := db.GetImageByID(id)
	if err != nil {
		return err
	}
	slog.Info("Checking image status...", "id", id)
	return render(r, w, generate.GalleryImage(image))
}
