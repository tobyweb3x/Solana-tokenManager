package http

import (
	"context"
	"net/http"
	frontend "tokenManager/client/pages"
)


func (a *App) indexPage(w http.ResponseWriter, r *http.Request) {
	if err := frontend.IndexPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func (a *App) showMintExtensions(w http.ResponseWriter, r *http.Request) {
	if err := frontend.MintExtensionsShowPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) showTokenExtensions(w http.ResponseWriter, r *http.Request) {
	if err := frontend.TokenExtensionsShowPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) mintExtensions(w http.ResponseWriter, r *http.Request) {
	if err := frontend.MintExtensions().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
