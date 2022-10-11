package client

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type header struct {
	key   string
	value string
}

type params struct {
	key   string
	value string
}

type client struct {
	url     string
	method  string
	body    io.Reader
	params  []*params
	headers []*header
	cookies []*http.Cookie
}

type response struct {
	Response *http.Response
	Error    error
}

func (c *client) Url(url string) *client {
	c.url = url
	return c
}

func (c *client) Header(key string, value string) *client {
	c.headers = append(c.headers, &header{
		key:   key,
		value: value,
	})

	return c
}

func (c *client) AddCookie(cookie *http.Cookie) *client {
	c.cookies = append(c.cookies, cookie)

	return c
}

func (c *client) Params(key string, value string) *client {
	c.params = append(c.params, &params{key: key, value: value})
	return c
}

func (c *client) Body(b []byte) *client {
	c.body = bytes.NewBuffer(b)
	return c
}

func (c *client) Exec() (*http.Response, error) {
	request, err := http.NewRequest(c.method, c.url, c.body)

	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	for _, item := range c.params {
		q.Add(item.key, item.value)
	}

	request.URL.RawQuery = q.Encode()

	request.Header.Add("access-control-allow-headers", "*")
	request.Header.Add("access-control-allow-origin", "*")
	request.Header.Add("accept", "application/json, text/plain, */*")
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("x-ratelimit-limit", "80")
	for _, item := range c.headers {
		request.Header.Add(item.key, item.value)
	}

	SERVER_READ_TIMEOUT, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	client := http.Client{
		Timeout: time.Second * time.Duration(SERVER_READ_TIMEOUT),
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *response) Decode(decode interface{}) error {
	data, err := ioutil.ReadAll(r.Response.Body)
	if err != nil {
		return err
	}

	if decode != nil {
		if err := json.Unmarshal(data, decode); err != nil {
			return err
		}
	}

	return nil
}

func New() *client {
	return &client{}
}

func Request(url string, method string) *client {
	return &client{
		url:    url,
		method: method,
	}
}

func (c *client) Get() *response {
	c.method = http.MethodGet
	res, err := c.Exec()
	return &response{res.Request.Response, err}
}

func (c *client) Post() *response {
	c.method = http.MethodPost
	res, err := c.Exec()
	return &response{res.Request.Response, err}
}

func (c *client) Patch() *response {
	c.method = http.MethodPatch
	res, err := c.Exec()
	return &response{res.Request.Response, err}
}

func (c *client) Delete() *response {
	c.method = http.MethodDelete
	res, err := c.Exec()
	return &response{res.Request.Response, err}
}

func (c *client) Put() *response {
	c.method = http.MethodPut
	res, err := c.Exec()
	return &response{res.Request.Response, err}
}
