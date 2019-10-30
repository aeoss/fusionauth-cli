package client

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

// Client is the HTTP Client that is authenticated using
// a Context.
type Client struct {
	context *Context
	client  *http.Client
}

// Context is the Client Context that will
type Context struct {
	URL string
	Key string
	JWT bool
}

// Init initializes a Client with the given Context.
func Init(context *Context) *Client {
	return &Client{
		context: context,
		client:  http.DefaultClient,
	}
}

// Request creates a request that holds the necessary config to
// nake an HTTP request to FusionAuth.
func (c *Client) Request(method string, path string, data []byte) (*http.Request, error) {
	fmt.Println(c.context)

	req, err := http.NewRequest(method, strings.Join([]string{c.context.URL, path}, "/"), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	var authKey = ""
	if c.context.JWT {
		authKey += "JWT "
	}

	authKey += c.context.Key

	req.Header.Add("Authorization", authKey)

	return req, nil
}

// Get performs a GET request using the given context.
func (c *Client) Get(path string) (*http.Response, error) {
	req, err := c.Request("GET", path, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Post performs a POST request using the given context.
func (c *Client) Post(path string, data []byte) (*http.Response, error) {
	req, err := c.Request("POST", path, data)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Patch performs a PATCH request using the given context.
func (c *Client) Patch(path string, data []byte) (*http.Response, error) {
	req, err := c.Request("PATCH", path, data)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Put performs a PUT request using the given context.
func (c *Client) Put(path string, data []byte) (*http.Response, error) {
	req, err := c.Request("PUT", path, data)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Delete performs a DELETE request using the given context.
func (c *Client) Delete(path string) (*http.Response, error) {
	req, err := c.Request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}
