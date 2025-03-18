package adapter

type Currency struct{}

func (c *Currency) Convert(quantity int, multiplier int) int {
	return quantity * multiplier
}
