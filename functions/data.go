package functions

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strings"
)

func MakeBodyFromData(data interface{}) (io.Reader, string, error) {
	switch x := data.(type) {
	case string:
		return strings.NewReader(x), "application/x-www-form-urlencoded", nil
	case []byte:
		return bytes.NewBuffer(x), "application/x-www-form-urlencoded", nil
	case io.Reader:
		return x, "", nil
	case nil:
		return nil, "", nil
	}
	if reflect.ValueOf(data).Kind() == reflect.Map {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewBuffer(b), "application/json", nil
	}
	if reflect.ValueOf(data).Kind() == reflect.Struct {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewBuffer(b), "application/json", nil
	}
	return nil, "", errors.New("invalid data type")
}
