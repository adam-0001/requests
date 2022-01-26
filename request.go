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

type session types.Session

var DefaultClientHello tls.ClientHelloID = tls.HelloChrome_Auto

func Client(timeout int, proxy string) (*session, error) {
	client, err := cclient.NewClient(tls.HelloChrome_Auto, proxy, true, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		return nil, err
	}
	return &session{&client, DefaultClientHello}, nil
}

func Session(timeout int, proxy string) (*session, error) {
	client, err := cclient.NewClient(tls.HelloChrome_Auto, proxy, true, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		return nil, err
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client.Jar = jar
	return &session{&client, DefaultClientHello}, nil
}

func (s *session) makeRequest(method string, url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
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
		functions.InferContentType(contentType, &newHeaders)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return types.Response{}, err
	}
	req.Header = newHeaders
	functions.FillNeededHeaders(host, &req.Header)
	start := time.Now()
	resp, err := s.Client.Do(req)

	duration := time.Since(start)
	if err != nil {
		return types.Response{}, err
	}
	defer resp.Body.Close()
	text, err := functions.Text(resp)
	if err != nil {
		return types.Response{}, err
	}
	return types.Response{HttpResponse: resp, Text: text, Elapsed: duration}, nil
}

func (s *session) Get(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.makeRequest(http.MethodGet, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Post(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.makeRequest(http.MethodPost, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Put(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.makeRequest(http.MethodPut, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Delete(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.makeRequest(http.MethodDelete, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Head(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.makeRequest(http.MethodHead, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Options(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.makeRequest(http.MethodOptions, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Trace(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.makeRequest(http.MethodTrace, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Patch(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.makeRequest(http.MethodPatch, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Connect(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.makeRequest(http.MethodConnect, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) SetProxy(proxy string) error {
	return cclient.SetProxy(s.Client, proxy, s.ClientHello)
}
