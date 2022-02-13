package requests

import (
	"encoding/json"
	"time"

	http "github.com/adam-0001/fhttp"
	tls "github.com/adam-0001/utls"
)

type Response struct {
	HttpResponse *http.Response
	statusCode   int
	status       string
	headers      map[string]string
	cookies      map[string]string
	encoding     string
	Elapsed      time.Duration
	Text         string
}

type Session struct {
	Client      *http.Client
	ClientHello tls.ClientHelloID
}

// func (r *Response) Text() (string, error) {
// 	if r.text == "" {
// 		responseBody, err := io.ReadAll(r.HttpResponse.Body)
// 		if err != nil {
// 			return "", err
// 		}
// 		r.text = string(responseBody)
// 		return r.text, nil
// 	}
// 	return r.text, nil
// }

func (r *Response) StatusCode() int {
	if r.statusCode == 0 {
		r.statusCode = r.HttpResponse.StatusCode
	}
	return r.statusCode
}

func (r *Response) Json(v interface{}) error {
	err := json.Unmarshal([]byte(r.Text), v)
	if err != nil {
		return err
	}
	return nil
}

// func (r *Response) JsonMap(v *map[string]string) error {
// 	err := json.Unmarshal([]byte(r.Text), &v)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (r *Response) Status() string {
	if r.status != "" {
		r.status = r.HttpResponse.Status
	}
	return r.status
}

func (r *Response) Headers() map[string]string {
	if r.headers == nil {
		r.headers = make(map[string]string)
		for k, v := range r.HttpResponse.Header {
			r.headers[k] = v[0]
		}
	}
	return r.headers
}

func (r *Response) Cookies() map[string]string {
	if r.cookies == nil {
		r.cookies = make(map[string]string)
		for _, v := range r.HttpResponse.Cookies() {
			r.cookies[v.Name] = v.Value
		}
	}
	return r.cookies
}

func (r *Response) Encoding() string {
	if r.encoding == "" {
		r.encoding = r.HttpResponse.Header.Get("Content-Type")
	}
	return r.encoding
}

// func (r *Response) Elapsed() time.Duration {
// 	return r.RElapsed
// }
