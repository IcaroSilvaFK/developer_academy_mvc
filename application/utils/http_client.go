package utils

import (
	"bytes"
	"context"
	"encoding/json"

	"io"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	client *http.Client
	ctx    context.Context
}

type HttpClientInterface interface {
	Get(url string, body interface{}, headers map[string]string) (*http.Response, error)
	CreateUrl(string, map[string]string) string
	Post(string, interface{}, interface{}) (*http.Response, error)
}

func NewHttpClient() HttpClientInterface {

	c := &http.Client{}

	return &HttpClient{
		client: c,
	}
}

func NewHttpClientWithContext(context context.Context, timeout time.Duration) HttpClientInterface {

	c := &http.Client{
		Timeout: timeout,
	}

	return &HttpClient{
		c, context,
	}
}

func (c *HttpClient) Get(url string, body interface{}, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)

	c.addDefaultHeaders(req)
	c.appendHeader(req, headers)

	if err != nil {
		return nil, err
	}

	r, err := c.client.Do(req)

	if err != nil {
		return r, err
	}
	defer r.Body.Close()

	err = c.toJson(r.Body, body)

	if err != nil {
		return r, err
	}

	return r, nil
}

func (c *HttpClient) Post(url string, body interface{}, result interface{}) (*http.Response, error) {

	bt, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bt))

	if err != nil {
		return nil, err
	}

	c.addDefaultHeaders(req)

	r, err := c.client.Do(req)

	if err != nil {
		return r, err
	}

	defer r.Body.Close()

	err = c.toJson(r.Body, result)

	return r, err
}

func (c *HttpClient) CreateUrl(base string, params map[string]string) string {

	if len(params) > 0 {
		for k, p := range params {
			if strings.Contains(base, "?") {
				base = base + "&" + k + "=" + p
			} else {
				base = base + "?" + k + "=" + p
			}
		}
	}
	return base
}

func (c *HttpClient) toJson(response io.ReadCloser, body interface{}) error {

	if body == nil {
		return nil
	}

	bt, err := io.ReadAll(response)

	if err != nil {
		return err
	}

	return json.Unmarshal(bt, body)
}

func (c *HttpClient) addDefaultHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func (c *HttpClient) appendHeader(req *http.Request, h map[string]string) {

	if len(h) == 0 {
		return
	}

	for k, v := range h {

		req.Header.Add(k, v)
	}

}
