package requests

//Note: Headers from this function will not maintain order
func MakeHeaders(h map[string]string) []map[string]string {
	headers := []map[string]string{}
	for key, value := range h {
		headers = append(headers, map[string]string{key: value})
	}
	return headers
}

