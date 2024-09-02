package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ujjwal405/url-shortner/pkg/store"
)

func TestShorten(t *testing.T) {
	testcases := []struct {
		name          string
		method        string
		body          userUrl
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{

		{

			name:   "invalidmethod",
			method: "GET",
			body:   userUrl{},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusMethodNotAllowed, recorder.Result().StatusCode)

			},
		},

		{

			name:   "invalidUrl",
			method: "POST",
			body: userUrl{
				Uri: "//www.example.com",
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
			},
		},

		{

			name:   "ok",
			method: "POST",
			body: userUrl{
				Uri: "https://www.youtube.com/watch?v=p7wDS1VV3n4&list=RDp7wDS1VV3n4&start_radio=1",
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Result().StatusCode)
			},
		},
	}

	for i := range testcases {
		tc := testcases[i]
		t.Run(tc.name, func(t *testing.T) {
			store := store.NewStore()
			handler := NewHandlers(store)
			res := httptest.NewRecorder()
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			req, err := http.NewRequest(tc.method, "/shorten", bytes.NewReader(data))
			require.NoError(t, err)
			handl := Make(handler.Shorten)
			handl.ServeHTTP(res, req)

			tc.checkResponse(t, res)
		})
	}
}

func TestShortCode(t *testing.T) {
	testcases := []struct {
		name          string
		method        string
		body          userUrl
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{

		{

			name:   "invalidmethod",
			method: "POST",
			body:   userUrl{},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusMethodNotAllowed, recorder.Result().StatusCode)

			},
		},

		{

			name:   "empty url",
			method: "GET",
			body: userUrl{
				Uri: "",
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
			},
		},
	}

	for i := range testcases {
		tc := testcases[i]
		t.Run(tc.name, func(t *testing.T) {
			store := store.NewStore()
			handler := NewHandlers(store)
			res := httptest.NewRecorder()
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			req, err := http.NewRequest(tc.method, "/shortCode", bytes.NewReader(data))
			require.NoError(t, err)
			handl := Make(handler.ShortCode)
			handl.ServeHTTP(res, req)

			tc.checkResponse(t, res)
		})
	}
}
