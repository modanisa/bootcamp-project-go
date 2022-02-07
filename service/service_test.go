package service

import (
	"bootcamp/mock"
	"bootcamp/model"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestServiceClientReturnQuotes(t *testing.T) {
	clientReturn := &model.ClientResponse{
		{Text: "text1", Author: "author1"},
		{Text: "text2", Author: "author2"},
		{Text: "text3", Author: "author3"},
		{Text: "text4", Author: "author1"},
		{Text: "text5", Author: "author2"},
		{Text: "text6", Author: "author1"},
	}
	client := mock.NewMockIClient(gomock.NewController(t))
	client.EXPECT().
		GetQuotes().
		Return(clientReturn, nil).
		Times(1)

	service := NewService(client)
	quotes, err := service.Quotes()

	expectedQuotes := &model.QuotesResponse{
		"author1": {"text1", "text4", "text6"},
		"author2": {"text2", "text5"},
		"author3": {"text3"},
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedQuotes, quotes)
}
func TestServiceClientReturnError(t *testing.T) {
	clientError := errors.New("test error")
	client := mock.NewMockIClient(gomock.NewController(t))
	client.EXPECT().
		GetQuotes().
		Return(nil, clientError).
		Times(1)

	service := NewService(client)
	response, err := service.Quotes()

	assert.ErrorIs(t, err, clientError)
	assert.Nil(t, response)
}
