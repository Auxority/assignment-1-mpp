package json_util

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"

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

func ParseInteger(value *string) (*int, error) {
	pattern := "[0-9]+"
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(*value, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("ParseInteger: the value '%s' is not an integer", *value)
	}

	lastNumber := matches[0]

	integer, err := strconv.Atoi(lastNumber)
	if err != nil {
		return nil, fmt.Errorf("ParseInteger: %w", err)
	}

	return &integer, nil
}

func ParseFloat(value *string) (*float64, error) {
	pattern := "[0-9]+.[0-9]+"
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(*value, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("ParseFloat: the value '%s' is not a float", *value)
	}

	lastNumber := matches[0]

	float, err := strconv.ParseFloat(lastNumber, 64)
	if err != nil {
		return nil, fmt.Errorf("ParseFloat: %w", err)
	}

	return &float, nil
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
