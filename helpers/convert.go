package helpers

import "encoding/json"

func Convert[T any](data any, destination T) (T, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return destination, err
	}
	err = json.Unmarshal(b, &destination)
	if err != nil {
		return destination, err
	}
	return destination, nil
}
