package service

import (
	"bootcamp/client"
	"bootcamp/model"
)

type IQuotesService interface {
	Quotes() (*model.QuotesResponse, error)
}

type QuotesService struct {
	Client client.IClient
}

func (s *QuotesService) Quotes() (*model.QuotesResponse, error) {
	quotes, err := s.Client.GetQuotes()
	if err != nil {
		return nil, err
	}

	m := model.QuotesResponse{}
	for _, v := range *quotes {
		if _, ok := m[v.Author]; !ok {
			m[v.Author] = []string{v.Text}
		} else {
			m[v.Author] = append(m[v.Author], v.Text)
		}
	}

	return &m, nil
}

func NewService(client client.IClient) IQuotesService {
	return &QuotesService{Client: client}
}
