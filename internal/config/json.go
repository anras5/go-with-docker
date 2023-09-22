package config

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// WriteJSON is used to write json data to the ResponseWrite
func WriteJSON(w http.ResponseWriter, status int, data interface{}, customHeaders ...http.Header) error {

	// marshal data
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// append custom headers
	if len(customHeaders) > 0 {
		for _, header := range customHeaders {
			for key, values := range header {
				for _, value := range values {
					w.Header().Add(key, value)
				}
			}
		}
	}

	// write to the Response Writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must contain only a single JSON value")
	}

	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	payload := struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}{
		Error:   true,
		Message: err.Error(),
	}

	return WriteJSON(w, statusCode, payload)
}
