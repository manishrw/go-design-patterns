# Singleton Pattern

## Description
The Singleton pattern ensures that a class has only one instance and provides a global point of access to it. In Go, this is typically implemented using package-level variables and sync.Once for thread safety.

## Use Cases
- Global configuration manager
- Logger
- Connection pool

## Go-Specific Considerations
- Use `sync.Once` to ensure thread-safe initialization
- Avoid using global state unless necessary
- Prefer dependency injection for testability

## Example
See `example.go` for a thread-safe Singleton implementation in Go.