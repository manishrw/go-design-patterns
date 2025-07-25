package abstractfactory

// Button is a UI component
 type Button interface {
	Paint() string
}

// Checkbox is a UI component
 type Checkbox interface {
	Check() string
}

// WinButton implements Button
 type WinButton struct{}

func (b *WinButton) Paint() string {
	return "Windows Button"
}

// MacButton implements Button
 type MacButton struct{}

func (b *MacButton) Paint() string {
	return "Mac Button"
}

// WinCheckbox implements Checkbox
 type WinCheckbox struct{}

func (c *WinCheckbox) Check() string {
	return "Windows Checkbox"
}

// MacCheckbox implements Checkbox
 type MacCheckbox struct{}

func (c *MacCheckbox) Check() string {
	return "Mac Checkbox"
}

// GUIFactory is the abstract factory
 type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// WinFactory creates Windows UI components
 type WinFactory struct{}

func (f *WinFactory) CreateButton() Button {
	return &WinButton{}
}
func (f *WinFactory) CreateCheckbox() Checkbox {
	return &WinCheckbox{}
}

// MacFactory creates Mac UI components
 type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
	return &MacButton{}
}
func (f *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}