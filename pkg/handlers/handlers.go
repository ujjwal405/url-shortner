package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ujjwal405/url-shortner/pkg/apierror"
	"github.com/ujjwal405/url-shortner/pkg/helper"
)

type handlers struct {
	memStore memStore
}

func NewHandlers(store memStore) *handlers {
	return &handlers{
		memStore: store,
	}
}

func (h *handlers) Shorten(w http.ResponseWriter, r *http.Request) error {
	var uri userUrl
	if r.Method != http.MethodPost {
		return apierror.InvalidMethod()
	}

	if err := json.NewDecoder(r.Body).Decode(&uri); err != nil {
		return apierror.InvalidJson()
	}

	u, _ := url.Parse(uri.Uri)
	if u.Scheme == "" || u.Host == "" {
		return apierror.InvalidJson()
	}

	shortCode := helper.GenerateShortCode([]byte(uri.Uri))

	h.memStore.InsertUrl(uri.Uri, shortCode)

	res := responseUrl{
		ShortUri: shortCode,
	}

	WriteJSON(w, http.StatusOK, res)
	return nil

}

func (h *handlers) ShortCode(w http.ResponseWriter, r *http.Request) error {
	var uri userUrl

	if r.Method != http.MethodGet {
		return apierror.InvalidMethod()
	}
	if err := json.NewDecoder(r.Body).Decode(&uri); err != nil {
		return apierror.InvalidJson()
	}

	if uri.Uri == "" {
		return apierror.NewAPIError(http.StatusBadRequest, fmt.Errorf(" provide the short url"))
	}

	longuri, err := h.memStore.GetUrl(uri.Uri)
	if err != nil {
		return err
	}

	Redirect(w, r, longuri)
	return nil
}
