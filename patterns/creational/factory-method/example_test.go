package factorymethod

import "testing"

func TestFactoryMethod(t *testing.T) {
	productA := FactoryMethod("A")
	if productA == nil || productA.Use() != "Using Product A" {
		t.Errorf("Expected Product A, got %v", productA)
	}

	productB := FactoryMethod("B")
	if productB == nil || productB.Use() != "Using Product B" {
		t.Errorf("Expected Product B, got %v", productB)
	}

	productUnknown := FactoryMethod("C")
	if productUnknown != nil {
		t.Errorf("Expected nil for unknown product type, got %v", productUnknown)
	}
}