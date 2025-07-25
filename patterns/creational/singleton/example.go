package singleton

import (
	"sync"
)

type singleton struct {
	// Add fields here
	Value int
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance returns the singleton instance
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}