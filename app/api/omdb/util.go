package omdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ReadJSONRequest(raw io.ReadCloser, data any) error {
	err := decodeJSON(raw, data)
	if err != nil {
		return fmt.Errorf("ReadJSONRequest: %w", err)
	}

	err = validateJSON(data)
	if err != nil {
		return fmt.Errorf("ReadJSONRequest: %w", err)
	}

	return nil
}

func WriteJSONResponse(writer http.ResponseWriter, data any) error {
	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		return fmt.Errorf("WriteJSONResponse: failed to encode JSON: %w", err)
	}

	return nil
}

func validateJSON(data any) error {
	err := validator.New(validator.WithRequiredStructEnabled()).Struct(data)
	if err != nil {
		return fmt.Errorf("validateJSON: failed to validate JSON: %w", err)
	}

	return nil
}

func decodeJSON(raw io.ReadCloser, data any) error {
	err := json.NewDecoder(raw).Decode(data)
	if err != nil {
		return fmt.Errorf("decodeJSON: failed to decode JSON: %w", err)
	}

	return nil
}
