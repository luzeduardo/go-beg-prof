package translation

import (
	"fmt"
	"log"
	"strings"
)

var __restTranslator = &RemoteService{}

// Remote service for external calls
type RemoteService struct {
	client HelloClient
	cache  map[string]string
}

type HelloClient interface {
	Translate(word, language string) (string, error)
}

func NewRemoteService(client HelloClient) *RemoteService {
	return &RemoteService{client: client,
		cache: make(map[string]string),
	}
}

func (s *RemoteService) Translate(word string, language string) (string, error) {
	word = strings.ToLower(word)
	language = strings.ToLower(language)

	key := fmt.Sprintf("%s:%s", word, language)
	tr, ok := s.cache[key]
	if ok {
		return tr, nil
	}

	resp, err := s.client.Translate(word, language)
	if err != nil {
		log.Println(err)
		return "", err
	}
	s.cache[key] = resp
	return resp, err
}
