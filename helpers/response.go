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

func GetHeaders(r *http.Response) map[string]string {
	headers := make(map[string]string)
	for k, v := range r.Header {
		headers[k] = v[0]
	}
	return headers
}

func GetCookies(r *http.Response) map[string]string {
	cookies := make(map[string]string)
	for _, v := range r.Cookies() {
		cookies[v.Name] = v.Value
	}
	return cookies
}
