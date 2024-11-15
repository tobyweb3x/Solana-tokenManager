package http

import (
	"context"
	"encoding/json"
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

	var (
		err       error
		inputData []string
		ok        bool
	)

	if err = r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if inputData, ok = r.Form["tokenInfo"]; len(inputData[0]) == 0 || !ok {
		http.Error(w, "post data is either empty or not in  the expected format", http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal([]byte(inputData[0]), &tokenInfo); err != nil {
		http.Error(w, fmt.Errorf("error unmarshaling tokenInfo: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	if tokenInfo.Wallet, err = solana.WalletFromPrivateKeyBase58(tokenInfo.MintAddressSecretKey); err != nil {
		http.Error(w, fmt.Errorf("wallet could not be created:%w", err).Error(), http.StatusBadRequest)
		return
	}

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

type TokenInfoData struct {
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

var tokenInfo TokenInfoData
