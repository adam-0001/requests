# Requests

**Requests is a simple http client inspired by [python requests](https://github.com/psf/requests). Built on top of [cclient](https://github.com/carcraftz/cclient) and [fhttp](https://github.com/carcraftz/fhttp), requests maintains simplicity while enabling features such as HTTP/2, ordered headers, and much more.**

---

# Quickstart

#

#### Get

```go
package main

import (
	"fmt"

	"github.com/adam-0001/requests"
)

func main() {
	//Headers will follow specified order
	headers := []map[string]string{
		{"hello": "world"},
	}
	resp, err := requests.Get("https://google.com", headers, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Text)
}

```
