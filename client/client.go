package client

import (
	"encoding/json"
	"net/http"

	"bootcamp/model"
)

type IClient interface {
	GetQuotes() (*model.ClientResponse, error)
}

type Client struct {
	url string
}

func (c *Client) GetQuotes() (*model.ClientResponse, error) {
	response, err := http.Get(c.url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	r := &model.ClientResponse{}
	err = decoder.Decode(r)

	return r, err
}

func NewClient(url string) IClient {
	return &Client{url: url}
}
