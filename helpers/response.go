package helpers

import (
	"io"

	http "github.com/adam-0001/fhttp"
)

func Text(r *http.Response) (string, []byte, error) {
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return "", []byte{}, err
	}
	return string(responseBody), responseBody, nil
}
