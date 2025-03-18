package adapter

type IAdapter interface {
	ConvertCurrency(quantity int, multiplier int) int
}
