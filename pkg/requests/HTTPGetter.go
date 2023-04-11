package requests

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type HttpGetter interface {
	Get(url string) (resp *http.Response, err error)
}

type DefaultHttpGetter struct {
}

func (d DefaultHttpGetter) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

var NotFoundErr = errors.New("Not found")
var BadRequestErr = errors.New("Bad request made")

func MakeApiRequest(h HttpGetter, url string) ([]byte, error) {
	resp, err := h.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, NotFoundErr
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}
