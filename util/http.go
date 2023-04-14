package util

import (
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	MethodGet    = "GET"
	MethodPOST   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
)

func HttpRequest(method, url string, f func(req *http.Request) (*http.Request, error)) (statuCode int, body []byte, err error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return -1, nil, errors.Wrap(err, "http new request error")
	}
	req, err = f(req)
	if err != nil {
		return -1, nil, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return -1, nil, errors.Wrap(err, "send request error")
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, errors.Wrap(err, "read response error")
	}
	return resp.StatusCode, body, nil
}