package service

import (
	"bootcamp/client"
	"fmt"
)

type IQuotesService interface {
	Quotes() error
}

type QuotesService struct {
	Client client.IClient
}

func (s *QuotesService) Quotes() error {
	quotes, err := s.Client.GetQuotes()
	if err != nil {
		return err
	}

	m := map[string][]string{}
	for _, v := range quotes {
		if _, ok := m[v.Author]; !ok {
			m[v.Author] = []string{v.Text}
		} else {
			m[v.Author] = append(m[v.Author], v.Text)
		}
	}

	fmt.Println(m)

	return err
}

func NewService(client client.IClient) IQuotesService {
	return &QuotesService{Client: client}
}
