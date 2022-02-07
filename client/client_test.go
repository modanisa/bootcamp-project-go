package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bootcamp/model"

	"github.com/stretchr/testify/assert"
)

func TestClientQuotes(t *testing.T) {
	serverReturn := `
	[
		{	"text":"text1",
			"author":"author1"
		},{
			"text":"text2",
			"author":"author2"
		},{
			"text":"text3",
			"author":"author1"
		}
	]`
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(serverReturn))
	}))
	defer testServer.Close()

	client := NewClient(testServer.URL)
	response, err := client.GetQuotes()

	expectedQuotes := &model.ClientResponse{
		{Text: "text1", Author: "author1"},
		{Text: "text2", Author: "author2"},
		{Text: "text3", Author: "author1"},
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedQuotes, response)
}

func TestClientEmptyQuotes(t *testing.T) {
	serverReturn := `[]`
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(serverReturn))
	}))
	defer testServer.Close()

	client := NewClient(testServer.URL)
	response, err := client.GetQuotes()

	emptyQuotes := &model.ClientResponse{}
	assert.Nil(t, err)
	assert.Equal(t, emptyQuotes, response)
}

func TestClientInvalidQuotes(t *testing.T) {
	serverReturn := `{}`
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(serverReturn))
	}))
	defer testServer.Close()

	client := NewClient(testServer.URL)
	_, err := client.GetQuotes()

	_, ok := err.(*json.UnmarshalTypeError)
	assert.True(t, ok)
}

func TestClientInvalidQuotesServerURL(t *testing.T) {
	client := NewClient("invalidServiceURL")

	_, err := client.GetQuotes()

	assert.NotNil(t, err)
}
