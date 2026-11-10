package kvstore

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("key not found")

type Store struct {
	data map[string][]byte
}

func New() *Store {
	m := Store{data: make(map[string][]byte)}
	return &m
}

func (s *Store) Set(key string, value []byte) {
	s.data[key] = value
}

func (s *Store) Get(key string) ([]byte, error) {
	value, ok := s.data[key]
	if !ok {
		return nil, fmt.Errorf("get %q: %w", key, ErrNotFound)
	}
	return value, nil
}
