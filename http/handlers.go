package http

import (
	"context"
	"net/http"
	"strconv"
	frontend "tokenManager/frontend/pages"
)

const (
	HxRequest = "HX-Request"
)

func (a *App) indexPage(w http.ResponseWriter, r *http.Request) {
	if err := frontend.IndexPage(nil).Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *App) showMintExtensions(w http.ResponseWriter, r *http.Request) {
	if isHXRequest(r) {
		if err := frontend.ShowMintExtensionsPartial().Render(context.Background(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if err := frontend.ShowMintExtensionsPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) showTokenExtensions(w http.ResponseWriter, r *http.Request) {
	if isHXRequest(r) {
		if err := frontend.ShowTokenExtensionsPartial().Render(context.Background(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if err := frontend.ShowTokenExtensionsPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *App) mintExtensions(w http.ResponseWriter, r *http.Request) {
	if isHXRequest(r) {
		if err := frontend.MintExtensionsPartial().Render(context.Background(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if err := frontend.MintExtensionsPage().Render(context.Background(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func isHXRequest(r *http.Request) bool {
	headerValue := r.Header.Get(HxRequest)
	if headerValue == "" {
		return false
	}

	if hxRequest, err := strconv.ParseBool(headerValue); (err == nil) && hxRequest {
		return true
	}

	return false
}
