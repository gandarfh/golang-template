package http

import "fmt"

func main() {
	req := struct {
		data string
	}{data: ""}

	err := Request.
		Get("localhost:5000").
		Header("Content-Type", "application/json").
		Header("Authentication", "Bearer <Token zika>").
		Decode(&req)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(req)
}

type header struct {
	key   string
	value string
}
type params struct {
	key   string
	value string
}
type cookie struct {
	key   string
	value string
}

type client struct {
	url     string
	params  []*params
	headers []*header
	body    []*any
}

type request struct{}

var Request = &request{}

func (c *client) Header(key string, value string) *client {
	return &client{}
}

func (c *client) AddCookie(key string, value string) *client {
	return &client{}
}

func (c *client) Params(key string, value string) *client {
	return &client{}
}

func (c *client) Body(value []byte) (error, *client) {
	return nil, &client{}
}

func (c *client) Decode(value interface{}) error {
	return nil
}

func (req *request) Post(url string) *client {

	return &client{
		url: url,
	}
}

func (req *request) Get(url string) *client {

	return &client{
		url: url,
	}
}

func (req *request) Patch(url string) *client {

	return &client{
		url: url,
	}
}

func (req *request) Delete(url string) *client {

	return &client{
		url: url,
	}
}

func (req *request) Put(url string) *client {

	return &client{
		url: url,
	}
}
