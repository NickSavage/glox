package interpreter

import (
	"fmt"
)

func (s *Storage) Assign(key string, value interface{}) error {
	_, ok := s.Memory[key]
	if !ok {
		if s.HasEnclosing {
			return s.Enclosing.Assign(key, value)
		}
		return fmt.Errorf("undefined variable: %s", key)
	}
	return s.Put(key, value)
}

func (s *Storage) Define(key string, value interface{}) error {
	return s.Put(key, value)

}

func (s *Storage) Put(key string, value interface{}) error {
	s.Memory[key] = value
	return nil
}

func (s *Storage) Get(key string) (interface{}, error) {
	value, ok := s.Memory[key]
	if !ok {
		if s.HasEnclosing {
			return s.Enclosing.Get(key)
		}
		return nil, fmt.Errorf("undefined variable: %s", key)
	}
	return value, nil
}
