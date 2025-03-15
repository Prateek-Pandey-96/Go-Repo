package strategy

type IStrategy interface {
	CalculateCharges(hours int) int
}
