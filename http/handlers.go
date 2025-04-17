package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	frontend "tokenManager/frontend/pages"

	"github.com/gagliardetto/solana-go"
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
	var err error

	if r.Method == http.MethodPost {
		if err = r.ParseForm(); err != nil {
			a.clientError(w, http.StatusBadRequest, fmt.Errorf("error parsing form request: %s", err.Error()))
			return
		}

		jsonStr := r.PostForm.Get("tokenInfo")
		if jsonStr == "" {
			a.clientError(w, http.StatusBadRequest, errors.New("invalid request body: form was empty"))
			return
		}

		if err = json.Unmarshal([]byte(jsonStr), &tokenInfo); err != nil {
			a.serverError(w, err)
			return
		}

		if tokenInfo.Wallet, err = solana.WalletFromPrivateKeyBase58(tokenInfo.MintAddressSecretKey); err != nil {
			a.serverError(w, err)
			return
		}

		fmt.Printf("\n\n%+v\n\n", tokenInfo)
	}

	if isHXRequest(r) {
		if err := frontend.ShowMintExtensionsPartial().Render(context.Background(), w); err != nil {
			a.serverError(w, err)
			return
		}
		return
	}

	if err := frontend.ShowMintExtensionsPage().Render(context.Background(), w); err != nil {
		a.serverError(w, err)
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

type TokenInfo struct {
	TokenStandard        string           `json:"tokenStandard"`
	TokenType            string           `json:"tokenType"`
	TokenName            string           `json:"tokenName"`
	TokenSymbol          string           `json:"tokenSymbol"`
	TokenUrl             string           `json:"tokenUrl"`
	MintAddressPublickey solana.PublicKey `json:"mintAddressPublickey"`
	MintAddressSecretKey string           `json:"mintAddressSecretKey"`
	Customized           bool             `json:"customized"`
	Wallet               *solana.Wallet
}

var tokenInfo TokenInfo
