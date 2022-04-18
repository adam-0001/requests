package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

var errInvalidType = errors.New("invalid data type")

func MakeBodyFromData(data interface{}) (io.Reader, error) {
	if data == nil {
		return nil, nil
	}
	// if reflect.ValueOf(data).Kind() == reflect.Struct {
	// 	v, err := json.Marshal(data)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return bytes.NewBuffer(v), nil
	// }
	// if x := fmt.Sprintf("%T", data); strings.HasPrefix(x, "map") {
	// 	v, err := json.Marshal(data)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return bytes.NewBuffer(v), nil
	// }
	switch x := data.(type) {
	case string:
		return bytes.NewBuffer([]byte(x)), nil
	case []byte:
		return bytes.NewBuffer(x), nil
	case io.Reader:
		return x, nil
		// case []interface{}:
		// 	v, err := json.Marshal(x)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	return bytes.NewBuffer(v), nil
		// }

	}
	res, err := json.Marshal(data)
	if err != nil {
		return nil, errInvalidType
	}
	return bytes.NewBuffer(res), nil
}
