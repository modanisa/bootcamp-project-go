package model

type ClientMiniResponse struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

type ClientResponse []ClientMiniResponse

type QuotesResponse map[string][]string
