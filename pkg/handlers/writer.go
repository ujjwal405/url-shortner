package handlers

import (
	"encoding/json"
	"log"

	"net/http"

	"github.com/ujjwal405/url-shortner/pkg/apierror"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func Redirect(w http.ResponseWriter, r *http.Request, longUrl string) {
	// w.Header().Set("Content-Type", "application/json")
	http.Redirect(w, r, longUrl, http.StatusMovedPermanently)
	//w.Header().Set("Location", longUrl)
	//w.WriteHeader(http.StatusFound)
}

func Make(h APIFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(apierror.APIError); ok {

				WriteJSON(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"StatusCode": http.StatusInternalServerError,
					"Msg":        "internal server error",
				}

				WriteJSON(w, http.StatusInternalServerError, errResp)

			}

		}
		log.Printf("request finished %v", r.Method)

	})
}
