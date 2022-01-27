package requests

import (
	"net/url"
	"time"

	"github.com/papermario8420/requests/functions"
	"github.com/papermario8420/requests/types"

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

func (s *session) MakeRequest(method string, url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
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
	return s.MakeRequest(http.MethodGet, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Post(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.MakeRequest(http.MethodPost, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Put(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.MakeRequest(http.MethodPut, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Delete(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.MakeRequest(http.MethodDelete, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Head(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.MakeRequest(http.MethodHead, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Options(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.MakeRequest(http.MethodOptions, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Trace(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.MakeRequest(http.MethodTrace, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Patch(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.MakeRequest(http.MethodPatch, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) Connect(url string, params map[string]string, headers []map[string]string, psuedoHeaderOrder []string, data interface{}, inferContentType bool) (types.Response, error) {
	return s.MakeRequest(http.MethodConnect, url, params, headers, psuedoHeaderOrder, data, inferContentType)
}

func (s *session) SetProxy(proxy string) error {
	return cclient.SetProxy(s.Client, proxy, s.ClientHello)
}

func (s *session) SetCookie(site url.URL, key string, value string) {
	cookie := &http.Cookie{
		Name:  key,
		Value: value,
	}
	s.Client.Jar.SetCookies(&site, []*http.Cookie{cookie})
}