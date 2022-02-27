package requests

import (
	"net/url"
	"time"

	"github.com/adam-0001/requests/functions"

	"github.com/adam-0001/cclient"
	http "github.com/adam-0001/fhttp"
	"github.com/adam-0001/fhttp/cookiejar"
	tls "github.com/adam-0001/utls"
)

var DefaultClientHello tls.ClientHelloID = tls.HelloChrome_Auto

func Client(timeout int, proxy string) (*Session, error) {
	client, err := cclient.NewClient(tls.HelloChrome_Auto, proxy, true, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		return nil, err
	}
	return &Session{&client, DefaultClientHello}, nil
}

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

func (s *Session) MakeRequest(method string, url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (Response, error) {
	url, host, err := functions.GetCompleteQuery(url, params)
	if err != nil {
		return Response{}, err
	}
	newHeaders := functions.MakeHeaders(headers, psuedoHeaderOrder)
	body, contentType, err := functions.MakeBodyFromData(data)
	if err != nil {
		return Response{}, err
	}
	if inferContentType {
		functions.InferContentType(contentType, &newHeaders)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return Response{}, err
	}
	req.Header = newHeaders
	functions.FillNeededHeaders(host, &req.Header)
	start := time.Now()
	//Defer a function to return an error if the request panics
	resp, err := s.Client.Do(req)
	duration := time.Since(start)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	text, err := functions.Text(resp)
	if err != nil {
		return Response{}, err
	}
	return Response{HttpResponse: resp, Text: text, Elapsed: duration}, nil
}

func (s *Session) Get(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (Response, error) {
	return s.MakeRequest(http.MethodGet, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *Session) Post(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (Response, error) {
	return s.MakeRequest(http.MethodPost, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *Session) Put(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (Response, error) {
	return s.MakeRequest(http.MethodPut, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *Session) Delete(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (Response, error) {
	return s.MakeRequest(http.MethodDelete, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *Session) Head(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (Response, error) {
	return s.MakeRequest(http.MethodHead, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *Session) Options(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (Response, error) {
	return s.MakeRequest(http.MethodOptions, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *Session) Trace(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (Response, error) {
	return s.MakeRequest(http.MethodTrace, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *Session) Patch(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (Response, error) {
	return s.MakeRequest(http.MethodPatch, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *Session) Connect(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (Response, error) {
	return s.MakeRequest(http.MethodConnect, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *Session) SetProxy(proxy string) error {
	return cclient.SetProxy(s.Client, proxy, s.ClientHello)
}

func (s *Session) SetTimeout(timeout int) {
	s.Client.Timeout = time.Duration(timeout) * time.Millisecond
}

func (s *Session) SetCookie(site url.URL, key string, value string) {
	cookie := &http.Cookie{
		Name:  key,
		Value: value,
	}
	s.Client.Jar.SetCookies(&site, []*http.Cookie{cookie})
}
