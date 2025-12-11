package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Envelope map[string]any

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func JSONResponse(w http.ResponseWriter, status int, data any, meta any) error {
	payload := Envelope{"data": data}
	if meta != nil {
		payload["meta"] = meta
	}
	return WriteJSON(w, status, payload)
}

func JSONError(w http.ResponseWriter, status int, errMsg string) {
	_ = WriteJSON(w, status, Envelope{"error": errMsg})
}

func ReadIDParam(r *http.Request) (int64, error) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		return 0, errors.New("missing id parameter")
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid id parameter: %w", err)
	}
	return id, nil
}

func ReadIntQuery(r *http.Request, key string, defaultVal int) (int, error) {
	q := r.URL.Query().Get(key)
	if q == "" {
		return defaultVal, nil
	}
	return strconv.Atoi(q)
}
