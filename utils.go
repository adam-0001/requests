package requests

import (
	"sync"

	"github.com/adam-0001/cclient"
)

// Note: Headers from this function will not maintain order
func MakeHeaders(h map[string]string) []map[string]string {
	headers := []map[string]string{}
	for key, value := range h {
		headers = append(headers, map[string]string{key: value})
	}
	return headers
}

// Create a new client with the same settiungs and cookiejar as the existing client
func NewClientFromSession(s *Session) (*Session, error) {
	newClient, err := cclient.NewClient(s.ClientHello, s.Proxy, true, s.Client.Timeout)
	if err != nil {
		return nil, err
	}

	// Share the cookie jar from the existing client
	newClient.Jar = s.Client.Jar

	// Return a new session with the new client
	return &Session{sync.Mutex{}, &newClient, s.ClientHello, s.Proxy}, nil
}
