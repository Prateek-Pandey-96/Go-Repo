package services

type IPendancyService interface {
	StartTracking(id int, tags []string) error
	StopTracking(id int) error
	GetCounts(tags []string) (int, error)
}
