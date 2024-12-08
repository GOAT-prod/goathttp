package json

import (
	"encoding/json"
	"github.com/GOAT-prod/goathttp/headers"
	"net/http"
)

func ReadRequest(r *http.Request, v any) error {
	if v == nil {
		return nil
	}

	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteResponse(w http.ResponseWriter, statusCode int, v any) error {
	if v == nil {
		return nil
	}

	w.Header().Set(headers.ContentTypeHeader(), headers.ContentTypeJSON())
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}
