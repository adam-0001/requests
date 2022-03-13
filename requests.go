package requests

import (
	"net/url"
	urllib "net/url"
	"strings"
	"time"

	"github.com/adam-0001/requests/functions"

	"github.com/adam-0001/cclient"
	http "github.com/adam-0001/fhttp"
	"github.com/adam-0001/fhttp/cookiejar"
	tls "github.com/adam-0001/utls"
)

var (
	defaultClientHello tls.ClientHelloID = tls.HelloChrome_Auto
	defaultClient, _                     = Client(30*time.Second, "")
)

func Client(timeout time.Duration, proxy string) (*Session, error) {
	client, err := cclient.NewClient(tls.HelloChrome_Auto, proxy, true, timeout)
	if err != nil {
		return nil, err
	}
	return &Session{&client, defaultClientHello}, nil
}

func NewSession(timeout time.Duration, proxy string) (*Session, error) {
	client, err := cclient.NewClient(tls.HelloChrome_Auto, proxy, true, timeout)
	if err != nil {
		return nil, err
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client.Jar = jar
	return &Session{&client, defaultClientHello}, nil
}

func (s *Session) MakeRequest(method string, url string, headers []map[string]string, data interface{}) (Response, error) {
	var resp Response
	parsedUrl, err := urllib.Parse(url)
	if err != nil {
		return resp, err
	}
	host := parsedUrl.Host
	newHeaders := functions.MakeHeaders(headers)
	body, _, err := functions.MakeBodyFromData(data)
	if err != nil {
		return resp, err
	}
	// if inferContentType {
	// 	functions.InferContentType(contentType, &newHeaders)
	// }
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return resp, err
	}
	req.Header = newHeaders
	functions.FillNeededHeaders(host, &req.Header)
	start := time.Now()
	//Defer a function to return an error if the request panics
	rawResp, err := s.Client.Do(req)
	duration := time.Since(start)
	if err != nil {
		return resp, err
	}
	defer rawResp.Body.Close()
	finalResp, bytes, err := functions.Text(rawResp)
	encoding, ok := rawResp.Header["Content-Encoding"]
	if ok {
		switch strings.ToLower(encoding[0]) {
		case "gzip":
			t, err := functions.UnGzip(bytes)
			if err == nil {
				finalResp = t
			}
		case "br":
			t, err := functions.UnBrotli(bytes)
			if err == nil {
				finalResp = t
			}
		case "deflate":
			t, err := functions.Enflate(bytes)
			if err == nil {
				finalResp = t
			}
		}
	}
	if err != nil {
		return resp, err
	}
	resp.HttpResponse = rawResp
	resp.Text = finalResp
	resp.Elapsed = duration
	return resp, nil
}

func (s *Session) Get(url string, headers []map[string]string, data interface{}) (Response, error) {
	return s.MakeRequest(http.MethodGet, url, headers, data)
}

func (s *Session) Post(url string, headers []map[string]string, data interface{}) (Response, error) {
	return s.MakeRequest(http.MethodPost, url, headers, data)
}

func (s *Session) Put(url string, headers []map[string]string, data interface{}) (Response, error) {
	return s.MakeRequest(http.MethodPut, url, headers, data)
}

func (s *Session) Delete(url string, headers []map[string]string, data interface{}) (Response, error) {
	return s.MakeRequest(http.MethodDelete, url, headers, data)
}

func (s *Session) Head(url string, headers []map[string]string, data interface{}) (Response, error) {
	return s.MakeRequest(http.MethodHead, url, headers, data)
}

func (s *Session) Options(url string, headers []map[string]string, data interface{}) (Response, error) {
	return s.MakeRequest(http.MethodOptions, url, headers, data)
}

func (s *Session) Trace(url string, headers []map[string]string, data interface{}) (Response, error) {
	return s.MakeRequest(http.MethodTrace, url, headers, data)
}

func (s *Session) Patch(url string, headers []map[string]string, data interface{}) (Response, error) {
	return s.MakeRequest(http.MethodPatch, url, headers, data)
}

func (s *Session) Connect(url string, headers []map[string]string, data interface{}) (Response, error) {
	return s.MakeRequest(http.MethodConnect, url, headers, data)
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

func Get(url string, headers []map[string]string, data interface{}) (Response, error) {
	return defaultClient.MakeRequest(http.MethodGet, url, headers, data)
}

func Post(url string, headers []map[string]string, data interface{}) (Response, error) {
	return defaultClient.MakeRequest(http.MethodPost, url, headers, data)
}

func Put(url string, headers []map[string]string, data interface{}) (Response, error) {
	return defaultClient.MakeRequest(http.MethodPut, url, headers, data)
}

func Delete(url string, headers []map[string]string, data interface{}) (Response, error) {
	return defaultClient.MakeRequest(http.MethodDelete, url, headers, data)
}

func Head(url string, headers []map[string]string, data interface{}) (Response, error) {
	return defaultClient.MakeRequest(http.MethodHead, url, headers, data)
}

func Options(url string, headers []map[string]string, data interface{}) (Response, error) {
	return defaultClient.MakeRequest(http.MethodOptions, url, headers, data)
}

func Trace(url string, headers []map[string]string, data interface{}) (Response, error) {
	return defaultClient.MakeRequest(http.MethodTrace, url, headers, data)
}

func Patch(url string, headers []map[string]string, data interface{}) (Response, error) {
	return defaultClient.MakeRequest(http.MethodPatch, url, headers, data)
}

func Connect(url string, headers []map[string]string, data interface{}) (Response, error) {
	return defaultClient.MakeRequest(http.MethodConnect, url, headers, data)
}
