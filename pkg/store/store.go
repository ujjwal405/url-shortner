package store

import (
	"sync"

	"github.com/ujjwal405/url-shortner/pkg/apierror"
)

type Store struct {
	lock sync.RWMutex
	urls map[string]string
}

func NewStore() *Store {
	return &Store{
		urls: make(map[string]string),
	}
}

func (s *Store) InsertUrl(url, code string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.urls[code] = url
}

func (s *Store) GetUrl(code string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	longURL, ok := s.urls[code]

	if !ok {
		return "", apierror.CodeNotExist()

	}

	return longURL, nil
}
