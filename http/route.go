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
	r.Get("/showTokenExtensions", a.showTokenExtensions)
	r.Get("/mintExtensions", a.mintExtensions)
	r.Post("/showMintExtensions", a.showMintExtensions)

	var staticFiles = fs.FS(frontend.AssetsDir)
	staticFs, _ := fs.Sub(staticFiles, "assets")
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(staticFs))))

	return r
}
