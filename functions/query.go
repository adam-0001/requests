package functions

import (
	urllib "net/url"
)

func GetCompleteQuery(url string, params map[string]string) (string, string, error) {
	parsedUrl, error := urllib.Parse(url)
	if error != nil {
		return "", "", error
	}
	for k, v := range params {
		query := parsedUrl.Query()
		query.Add(k, v)
		parsedUrl.RawQuery = query.Encode()
	}
	return parsedUrl.String(), parsedUrl.Host, nil
}

func ToQueryString(i map[string]string) string {
	query := urllib.Values{}
	for k, v := range i {
		query.Add(k, v)
	}
	return query.Encode()
}
