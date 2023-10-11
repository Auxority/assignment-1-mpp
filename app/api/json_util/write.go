package json_util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSONResponse(writer http.ResponseWriter, data any) error {
	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		return fmt.Errorf("WriteJSONResponse: failed to encode JSON: %w", err)
	}

	return nil
}
