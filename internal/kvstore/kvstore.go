package kvstore

import (
	"bytes"
	"errors"
	"fmt"
	"maps"
	"slices"
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
	clone := bytes.Clone(value)
	s.data[key] = clone
}

func (s *Store) Get(key string) ([]byte, error) {
	value, ok := s.data[key]
	if !ok {
		return nil, fmt.Errorf("get %q: %w", key, ErrNotFound)
	}
	clone := bytes.Clone(value)
	return clone, nil
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

func (s *Store) Keys() []string {
	return slices.Sorted(maps.Keys(s.data))
}
