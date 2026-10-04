package request

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type envelope struct {
	Success bool 	`json:"success"`
	Data any 		`json:"data"`
	Error string 	`json:"error,omitempty"`
}

func Success(w http.ResponseWriter, status int, data any) {
	Write(w, status, envelope{
		Success: true,
		Data: data,
	})
}

func Error(w http.ResponseWriter, status int, err error) {
	message := "Internal server error"

	if status < http.StatusInternalServerError && err != nil {
		message = err.Error()
	}

	Write(w, status, envelope{
		Success: false,
		Data: nil,
		Error: message,
	})
}

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Read(r *http.Request, data any) error {
	r.Body = http.MaxBytesReader(nil, r.Body, 1<<20)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	var extra any
	err = decoder.Decode(&extra);
	if !errors.Is(err, io.EOF) {
		if err == nil {
			return errors.New("request body must contain one JSON object")
		}
		
		return err
	}
	
	return nil
}