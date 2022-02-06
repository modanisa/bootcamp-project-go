package client

import (
	"bootcamp/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IClient interface {
	GetQuotes() (model.ClientResponse, error)
}

type Client struct {
	url string
}

func (c *Client) GetQuotes() (model.ClientResponse, error) {
	response, err := http.Get(c.url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	read, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	r := &model.ClientResponse{}
	err = json.Unmarshal(read, r)
	return *r, err
}

func NewClient(url string) IClient {
	return &Client{url: url}
}
