package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sentinel/med_location/types"
)

func WriteResponse(w http.ResponseWriter, data any) error {
	// WriteResponse writes the encoded data as a response using the provided http.ResponseWriter.

	// encode data
	jsonData, err := EncodeRequest(data)
	if err != nil {
		return err
	}
	// write response
	_, err = fmt.Fprintf(w, "data: %s\n\n", jsonData)
	return err
	// return err
}
func InternalServerError(w http.ResponseWriter, data any) {
	// InternalServerError is a function that takes in an http.ResponseWriter and a data parameter.
	// It encodes the data using the 'EncodeRequest' function, sets the HTTP status code to 500 (Internal Server Error),
	// and writes the encoded data to the response.
	//
	// Parameters:
	// - w (http.ResponseWriter): the response writer to write the encoded data to
	// - data (any): the data to be encoded and written to the response
	//
	// Returns:
	// - None

	b, _ := EncodeRequest(data)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(b)
}

// encode request data to json bytes
func EncodeRequest(data any) ([]byte, error) {
	// EncodeRequest encodes the given data into JSON format.

	// Parameters:
	// - data (any): The data to be encoded into JSON format.

	// Returns:
	// - ([]byte, error): The encoded data as a byte slice and any error that occurred during the encoding process.

	return json.Marshal(data)
}

// DecodeRequest decodes data from an io.Reader and assigns it to a types.Message pointer.
// If there is an error during the decoding process, it returns the error.
func DecodeRequest(r io.Reader, msg *types.Message) error {
	if err := json.NewDecoder(r).Decode(msg); err != nil {
		return err
	}
	return nil
}
