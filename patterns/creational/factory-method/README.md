# Factory Method Pattern

## Description
The Factory Method pattern defines an interface for creating an object, but lets subclasses alter the type of objects that will be created. In Go, this is typically implemented using interfaces and constructor functions.

## Use Cases
- Database connector factory
- Logger factory
- Notification sender factory

## Go-Specific Considerations
- Use interfaces to define product types
- Use constructor functions for flexibility
- Prefer composition over inheritance

## Example
See `example.go` for a Go implementation of the Factory Method pattern.