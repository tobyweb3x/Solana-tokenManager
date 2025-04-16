package http

import (
	"io/fs"

	frontend "tokenManager/frontend"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
)

func (a *App) Routes() *chi.Mux {
	r := chi.NewRouter()

	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))
	r.NotFoundHandler()
	r.MethodNotAllowedHandler()
	// r.MethodNotAllowed()
	// r.NotFound()

	r.Get("/", a.indexPage)
	r.Get("/mintExtensions", a.mintExtensions)
	r.Get("/showTokenExtensions", a.showTokenExtensions)
	r.Post("/showTokenExtensions", a.showTokenExtensions)
	r.Get("/showMintExtensions", a.showMintExtensions)
	r.Post("/showMintExtensions", a.showMintExtensions)

	var staticFiles = fs.FS(frontend.AssetsDir)
	staticFs, _ := fs.Sub(staticFiles, "public")
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.FS(staticFs))))
	return r
}
