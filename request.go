package requests

import (
	"time"

	"requests/functions"
	"requests/types"

	"github.com/papermario8420/cclient"
	http "github.com/papermario8420/fhttp"
	"github.com/papermario8420/fhttp/cookiejar"
	tls "github.com/papermario8420/utls"
)

type Session types.Session

var DefaultClientHello tls.ClientHelloID = tls.HelloChrome_Auto

func NewSession(timeout int, proxy string) (*Session, error) {
	client, err := cclient.NewClient(tls.HelloChrome_Auto, proxy, true, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		return nil, err
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client.Jar = jar
	return &Session{&client, DefaultClientHello}, nil
}

func (s *Session) MakeRequest(method string, url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	url, host, err := functions.GetCompleteQuery(url, params)
	if err != nil {
		return types.Response{}, err
	}
	newHeaders := functions.MakeHeaders(headers, psuedoHeaderOrder)
	body, contentType, err := functions.MakeBodyFromData(data)
	if err != nil {
		return types.Response{}, err
	}
	if inferContentType {
		_, ok := newHeaders["Content-Type"]
		_, ok1 := newHeaders["content-type"]
		if !ok && !ok1 {
			newHeaders["Content-Type"] = []string{contentType}
			newHeaders["Header-Order:"] = append(newHeaders["Header-Order:"], "content-type")
		}
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return types.Response{}, err
	}
	if len(headers) > 0 {
		req.Header = newHeaders
		functions.FillNeededHeaders(host, &req.Header)
	} else {
		functions.FillNeededHeaders(host, &req.Header)
	}
	start := time.Now()
	resp, err := s.Client.Do(req)
	duration := time.Since(start)
	if err != nil {
		return types.Response{}, err
	}
	return types.Response{HttpResponse: resp, Elapsed: duration}, nil
}
