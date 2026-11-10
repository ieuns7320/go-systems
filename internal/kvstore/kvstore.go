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

func (s *Store) Delete(key string) error {
	_, ok := s.data[key]

	if !ok {
		return fmt.Errorf("delete %q: %w", key, ErrNotFound)
	}
	delete(s.data, key)
	return nil
}

func (s *Store) Len() int {
	return len(s.data)
}
