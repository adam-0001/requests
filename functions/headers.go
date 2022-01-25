package functions

import "strings"

func MakeHeaders(headers []map[string]string, psuedoHeaderOrder []string) map[string][]string {
	var reqHeaders = map[string][]string{
		"Header-Order:":  {},
		"PHeader-Order:": {},
	}
	for _, v := range headers {
		for key, val := range v {
			reqHeaders[key] = []string{val}
			reqHeaders["Header-Order:"] = append(reqHeaders["Header-Order:"], strings.ToLower(key))
		}
	}
	reqHeaders["PHeader-Order:"] = append(reqHeaders["PHeader-Order:"], psuedoHeaderOrder...)
	return reqHeaders
}
