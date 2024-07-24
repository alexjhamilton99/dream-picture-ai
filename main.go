package main

import (
	"dream-picture-ai/handler"
	"dream-picture-ai/pkg/sb"
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()
	router.Use(handler.WithUser)

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(handler.HandleHomeIndex))
	router.Get("/login", handler.Make(handler.HandleLoginIndex))
	router.Get("/sign-up", handler.Make(handler.HandleSignUpIndex))
	router.Post("/login", handler.Make(handler.HandleLoginCreate))
	router.Post("/logout", handler.Make(handler.HandleLogoutCreate))
	router.Post("/sign-up", handler.Make(handler.HandleSignUpCreate))
	// router.Get("/auth/callback", handler.Make(handler.HandleAuthCallback))

	router.Group(func(auth chi.Router) {
		auth.Use(handler.WithAuth)
		auth.Get("/settings", handler.Make(handler.HandleSettingsIndex))
	})

	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("Application running...", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return sb.Init()
}
