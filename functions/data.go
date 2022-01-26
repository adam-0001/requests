package functions

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"reflect"
)

var InvalidTypeError = errors.New("Invalid Data Type passed to Data")
var InvalidStructError = errors.New("Structs Cannot Be passed to Data")

func MakeBodyFromData(data interface{}) (io.Reader, string, error) {
	switch x := data.(type) {
	case nil:
		return nil, "", nil
	case string:
		return bytes.NewBuffer([]byte(x)), "application/x-www-form-urlencoded", nil
	case []byte:
		return bytes.NewBuffer(x), "application/x-www-form-urlencoded", nil
	case io.Reader:
		return x, "", nil
	case map[string]string:
		v, err := json.Marshal(x)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewBuffer(v), "application/json", nil
	case map[string]interface{}:
		v, err := json.Marshal(x)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewBuffer(v), "application/json", nil
	case map[string][]string:
		v, err := json.Marshal(x)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewBuffer(v), "application/json", nil
	case map[string][]interface{}:
		v, err := json.Marshal(x)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewBuffer(v), "application/json", nil
	case map[string]int:
		v, err := json.Marshal(x)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewBuffer(v), "application/json", nil
	case map[string]float64:
		v, err := json.Marshal(x)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewBuffer(v), "application/json", nil
	case map[string]bool:
		v, err := json.Marshal(x)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewBuffer(v), "application/json", nil
	case []interface{}:
		v, err := json.Marshal(x)
		if err != nil {
			return nil, "", err
		}
		return bytes.NewBuffer(v), "application/json", nil
	}
	if reflect.ValueOf(data).Kind() == reflect.Struct {
		return nil, "", InvalidStructError
	}
	return nil, "", InvalidTypeError
}
