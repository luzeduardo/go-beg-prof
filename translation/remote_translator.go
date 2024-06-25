package translation

import (
	"strings"
)

var __rest.Translator = &RemoteSRemoteService{}

//Remote service for external calls
type RemoteService struct {
	client HelloClient
}

type HelloClient interface {
	Translate(word, language string) (string, error)
}

func NewRemoteService(client HelloClient) *RemoteService {
	return &RemoteService{ client: client }
}

func (s *RemoteService) Translate(word string, language string) (string, error) {
	word = strings.ToLower(word)
	language = strings.ToLower(language)
	resp, _ := s.client.Translate(word, language)
	return resp
}