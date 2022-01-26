package functions

import (
	"io"

	http "github.com/papermario8420/fhttp"
)

func Text(r *http.Response) (string, error) {
	// if r.text == "" {
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(responseBody), nil
	// }
	// return r.text, nil
}
