package omdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ReadJSONRequest(raw io.ReadCloser, data any) error {
	err := json.NewDecoder(raw).Decode(data)
	if err != nil {
		return fmt.Errorf("ReadJSONRequest: failed to decode JSON: %w", err)
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
