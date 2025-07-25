# Builder Pattern

## Description
The Builder pattern separates the construction of a complex object from its representation, allowing the same construction process to create different representations. In Go, this is often implemented using method chaining.

## Use Cases
- Query builder
- Configuration builder
- Complex object construction (e.g., HTTP requests)

## Go-Specific Considerations
- Use method chaining for fluent APIs
- Use value or pointer receivers as appropriate
- Keep builders immutable if possible

## Example
See `example.go` for a Go implementation of the Builder pattern.