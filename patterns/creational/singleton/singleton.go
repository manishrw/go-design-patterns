package singleton

import (
	"sync"
)

// Singleton represents a thread-safe singleton instance
type Singleton struct {
	value any
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance returns the singleton instance
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

// SetValue sets a new value
func (s *Singleton) SetValue(v interface{}) {
	s.value = v
}

// GetValue returns the stored value
func (s *Singleton) GetValue() interface{} {
	return s.value
}
