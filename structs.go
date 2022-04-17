package requests

import (
	"encoding/json"
	"time"

	http "github.com/adam-0001/fhttp"
	tls "github.com/adam-0001/utls"
)

type Response struct {
	HttpResponse    *http.Response
	statusCode      int
	headers         map[string]string
	cookies         map[string]string
	encoding        string
	Elapsed         time.Duration
	Text            string
	RedirectHistory []string
}

type Session struct {
	Client      *http.Client
	ClientHello tls.ClientHelloID
}

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

func (r *Response) Url() string {
	return r.HttpResponse.Request.URL.String()
}

func (r *Response) Status() string {
	return r.HttpResponse.Status
}

func (r *Response) Headers() map[string]string {
	if len(r.headers) == 0 {
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
