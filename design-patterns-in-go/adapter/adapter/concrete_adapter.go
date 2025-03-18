package adapter

type ConcreteAdapter struct {
	Adaptee Currency
}

func (ca *ConcreteAdapter) GetAdapter(adaptee Currency) *ConcreteAdapter {
	ca.Adaptee = adaptee
	return ca
}

func (ca *ConcreteAdapter) ConvertCurrency(quantity int, multiplier int) int {
	return ca.Adaptee.Convert(quantity, multiplier/2) + 2
}
