package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // 1MB

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(data); err != nil {
		return err
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must have only a single json value")
	}

	return nil
}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, header http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if header != nil {
		for key, value := range header {
			w.Header()[key] = value
		}
	}

	if len(header) > 0 {
		for key, value := range header {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, status int) error {
	var statusCode int

	if status == 0 {
		statusCode = http.StatusBadRequest
	}

	payload := JSONResponse{
		Error:   true,
		Message: err.Error(),
	}

	return app.writeJSON(w, statusCode, payload, nil)
}
