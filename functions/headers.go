package functions

import (
	"strings"

	http "github.com/papermario8420/fhttp"
)

func MakeHeaders(headers []map[string]string, psuedoHeaderOrder []string) map[string][]string {
	var reqHeaders = map[string][]string{
		"Header-Order:": {},
		// "PHeader-Order:": {},
	}
	for _, v := range headers {

		for key, val := range v {
			reqHeaders[key] = []string{val}
			reqHeaders["Header-Order:"] = append(reqHeaders["Header-Order:"], strings.ToLower(key))
		}
	}
	if len(psuedoHeaderOrder) > 0 {
		reqHeaders["PHeader-Order:"] = append(reqHeaders["PHeader-Order:"], psuedoHeaderOrder...)
	}
	return reqHeaders
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func FillNeededHeaders(host string, headers *http.Header) {
	defHeaders := DefaultHeaders(host)
	keysLower := make([]string, len(*headers))
	i := 0
	for k := range *headers {
		keysLower[i] = strings.ToLower(k)
		i++
	}
	for k := range defHeaders {
		lower := strings.ToLower(k)
		if !contains(keysLower, lower) {
			(*headers)[k] = defHeaders[k]
			(*headers)["Header-Order:"] = append((*headers)["Header-Order:"], lower)
		}
	}
	// for k, v := range defHeaders {
	// 	key := strings.ToLower(k)

	// 	if _, ok := (*headers)[key]; !ok {
	// 		// fmt.Println("Edited:", key, ok)
	// 		(*headers)[k] = v
	// 		(*headers)["Header-Order:"] = append((*headers)["Header-Order:"], key)
	// 	}
	// }
}

func DefaultHeaders(host string) map[string][]string {
	return (map[string][]string{
		"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"},
		"Accept":          {"*/*"},
		"Accept-Encoding": {"gzip, deflate, br, utf-8"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"Connection":      {"keep-alive"},
		"Host":            {host},
	})
}
