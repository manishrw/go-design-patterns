# Contributing to Go Design Patterns

## How to Contribute
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/pattern-name`)
3. Commit your changes (`git commit -am 'Add Pattern: pattern-name'`)
4. Push to the branch (`git push origin feature/pattern-name`)
5. Create a Pull Request

## Pattern Implementation Guidelines
Each pattern should include:
- README.md with pattern description
- example.go with implementation
- example_test.go with tests

## Code Style
- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use meaningful variable and function names
- Add comments for complex logic
- Include examples in tests

## Testing Requirements
- Write unit tests for all implementations
- Include examples in test files
- Ensure thread safety where applicable
- Add benchmarks for performance-critical code

## Documentation Guidelines
- Clear pattern description
- Problem it solves
- When to use it
- Implementation considerations
- Real-world examples
- Best practices 