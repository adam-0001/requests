package helpers

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io/ioutil"

	"github.com/andybalholm/brotli"
)

func UnGzip(bd []byte) (string, error) {
	gz, _ := gzip.NewReader(bytes.NewReader(bd))
	defer gz.Close()
	body, err := ioutil.ReadAll(gz)
	return string(body), err
}

func UnBrotli(bd []byte) (string, error) {
	br := brotli.NewReader(bytes.NewReader(bd))
	body, err := ioutil.ReadAll(br)
	return string(body), err
}

func Enflate(bd []byte) (string, error) {
	zr, _ := zlib.NewReader(bytes.NewReader(bd))
	defer zr.Close()
	body, err := ioutil.ReadAll(zr)
	return string(body), err
}
