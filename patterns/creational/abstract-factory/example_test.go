package abstractfactory

import "testing"

func TestWinFactory(t *testing.T) {
	factory := &WinFactory{}
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()
	if button.Paint() != "Windows Button" {
		t.Errorf("Expected Windows Button, got %s", button.Paint())
	}
	if checkbox.Check() != "Windows Checkbox" {
		t.Errorf("Expected Windows Checkbox, got %s", checkbox.Check())
	}
}

func TestMacFactory(t *testing.T) {
	factory := &MacFactory{}
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()
	if button.Paint() != "Mac Button" {
		t.Errorf("Expected Mac Button, got %s", button.Paint())
	}
	if checkbox.Check() != "Mac Checkbox" {
		t.Errorf("Expected Mac Checkbox, got %s", checkbox.Check())
	}
}