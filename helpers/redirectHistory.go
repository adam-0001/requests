package helpers

import (
	http "github.com/adam-0001/fhttp"
)

func SetRedirectUrlHistory(res *http.Response) (h []string) {
	for res != nil {
		req := res.Request
		h = append(h, req.URL.String())
		res = req.Response
	}
	for l, r := 0, len(h)-1; l < r; l, r = l+1, r-1 {
		h[l], h[r] = h[r], h[l]
	}
	return
}
