package strategy

type NormalStrategy struct{}
type PeakHoursStrategy struct{}

func (ns *NormalStrategy) CalculateCharges(hours int) int {
	return 10 * hours
}

func (ps *PeakHoursStrategy) CalculateCharges(hours int) int {
	return 12*hours + 40
}
