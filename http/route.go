package http

import (
	"io/fs"

	frontend "tokenManager/client"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
)

func (a *App) Routes() *chi.Mux {
	r := chi.NewRouter()

	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Get("/", a.indexPage)
	r.Get("/showMintExtensions", a.showMintExtensions)
	r.Get("/showTokenExtensions", a.showTokenExtensions)
	r.Get("/mintExtensions", a.mintExtensions)


	var staticFiles = fs.FS(frontend.AssetsDir)
	staticFs, _ := fs.Sub(staticFiles, "assets")
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(staticFs))))

	return r
}
