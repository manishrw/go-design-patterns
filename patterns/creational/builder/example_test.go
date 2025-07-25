package builder

import "testing"

func TestUserBuilder(t *testing.T) {
	builder := NewUserBuilder()
	user := builder.Name("Alice").Email("alice@example.com").Age(30).Build()

	if user.Name != "Alice" {
		t.Errorf("Expected Name to be Alice, got %s", user.Name)
	}
	if user.Email != "alice@example.com" {
		t.Errorf("Expected Email to be alice@example.com, got %s", user.Email)
	}
	if user.Age != 30 {
		t.Errorf("Expected Age to be 30, got %d", user.Age)
	}
}