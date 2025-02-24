package utils

import "fmt"

type set[K comparable, V any] struct {
	m map[K]V
}

func newSet[K comparable, V any]() *set[K, V] {
	return &set[K, V]{
		m: make(map[K]V),
	}
}

func (s *set[K, V]) verifyUnique(key K) error {
	if _, exists := s.m[key]; exists {
		return fmt.Errorf("key %v already exists", key)
	}
	return nil
}

func (s *set[K, V]) add(key K, val V) {
	s.m[key] = val
}

func (s *set[K, V]) remove(key K) {
	delete(s.m, key)
}

func (s *set[K, V]) hasKey(key K) bool {
	_, exists := s.m[key]
	return exists
}

func (s *set[K, V]) modifyKey(oldKey, newKey K, val V) {
	s.remove(oldKey)
	s.add(newKey, val)
}

func (s *set[K, V]) get(key K) (V, error) {
	val, exists := s.m[key]
	if !exists {
		return val, fmt.Errorf("key %v does not exists", key)
	}
	return val, nil
}

func (s *set[K, V]) getKeys() []K {
	count := len(s.m)
	keys := make([]K, count)
	i := 0
	for k := range s.m {
		keys[i] = k
		i++
	}
	return keys
}

func (s *set[K, V]) size() int {
	return len(s.m)
}

func (s *set[K, V]) clear() {
	s.m = make(map[K]V)
}

func (s *set[K, V]) entries() map[K]V {
	return s.m
}

func (s *set[K, V]) getValues() []V {
	count := len(s.m)
	values := make([]V, count)
	i := 0
	for _, v := range s.m {
		values[i] = v
		i++
	}
	return values
}
