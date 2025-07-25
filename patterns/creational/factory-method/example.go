package factorymethod

// Product is the interface for products
 type Product interface {
	Use() string
}

// ConcreteProductA implements Product
 type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Using Product A"
}

// ConcreteProductB implements Product
 type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Using Product B"
}

// FactoryMethod creates a Product based on type
func FactoryMethod(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}