package session

import (
	"io/ioutil"
	"net/http"
)

const ()

// Reference: go-github/github/github.go

// Client http client
type Client struct {
	httpClient *http.Client
}

// NewClient factory Client struct
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		httpClient: httpClient,
	}
}

// Get exec HTTP GET
func (c *Client) Get(url string, header map[string]string) ([]byte, error) {
	req, err := c.newRequest("GET", url, header)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// Do GET only now
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) newRequest(method string, url string, header map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	return req, nil
}
