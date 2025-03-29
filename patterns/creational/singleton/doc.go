// Package singleton demonstrates the Singleton design pattern in Go.
//
// The Singleton pattern ensures a class has only one instance and provides a global
// point of access to it. This implementation is thread-safe using sync.Once.
//
// Example usage:
//
//	instance := singleton.GetInstance()
//	instance.IncrementCounter()
//	count := instance.GetCounter()
//
// See the README.md for more details on the Singleton pattern.
package singleton
