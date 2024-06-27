package translation

import (
	"strings"
)

var __restTranslator = &RemoteService{}

// Remote service for external calls
type RemoteService struct {
	client HelloClient
}

type HelloClient interface {
	Translate(word, language string) (string, error)
}

func NewRemoteService(client HelloClient) *RemoteService {
	return &RemoteService{client: client}
}

func (s *RemoteService) Translate(word string, language string) (string, error) {
	word = strings.ToLower(word)
	language = strings.ToLower(language)
	resp, err := s.client.Translate(word, language)
	return resp, err
}
