package utils

import "fmt"

type Set[K comparable, V any] struct {
	m map[K]V
}

func NewSet[K comparable, V any]() *Set[K, V] {
	return &Set[K, V]{
		m: make(map[K]V),
	}
}

func (s *Set[K, V]) VerifyUnique(key K) error {
	if _, exists := s.m[key]; exists {
		return fmt.Errorf("key %v already exists", key)
	}
	return nil
}

func (s *Set[K, V]) Add(key K, val V) {
	s.m[key] = val
}

func (s *Set[K, V]) Remove(key K) {
	delete(s.m, key)
}

func (s *Set[K, V]) HasKey(key K) bool {
	_, exists := s.m[key]
	return exists
}

func (s *Set[K, V]) EditKey(oldKey, newKey K, val V) {
	s.Remove(oldKey)
	s.Add(newKey, val)
}

func (s *Set[K, V]) Get(key K) (V, error) {
	val, exists := s.m[key]
	if !exists {
		return val, fmt.Errorf("key %v does not exists", key)
	}
	return val, nil
}

func (s *Set[K, V]) GetKeys() []K {
	count := len(s.m)
	keys := make([]K, count)
	i := 0
	for k := range s.m {
		keys[i] = k
		i++
	}
	return keys
}

func (s *Set[K, V]) Size() int {
	return len(s.m)
}

func (s *Set[K, V]) Clear() {
	s.m = make(map[K]V)
}

func (s *Set[K, V]) Entries() map[K]V {
	return s.m
}

func (s *Set[K, V]) GetValues() []V {
	count := len(s.m)
	values := make([]V, count)
	i := 0
	for _, v := range s.m {
		values[i] = v
		i++
	}
	return values
}
