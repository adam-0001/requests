package requests

import (
	"encoding/json"
	"os"
	"time"

	http "github.com/adam-0001/fhttp"
	tls "github.com/adam-0001/utls"
)

type Response struct {
	HttpResponse    *http.Response
	Elapsed         time.Duration
	Text            string
	RedirectHistory []string
	Headers         map[string]string
	Status          string
	Url             string
	Cookies         map[string]string
	Encoding        string
	StatusCode      int
}

type Session struct {
	Client      *http.Client
	ClientHello tls.ClientHelloID
}

func (r *Response) Json(v interface{}) error {
	err := json.Unmarshal([]byte(r.Text), v)
	if err != nil {
		return err
	}
	return nil
}

func (r *Response) SaveToFile(filePath string) error {
	return os.WriteFile(filePath, []byte(r.Text), 0644)
}
