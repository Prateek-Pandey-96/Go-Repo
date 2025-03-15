package main

import "github.com/prateek96/pendancySystem/services"

type PendancySystem struct {
	pendancyService services.IPendancyService
}

func (ps *PendancySystem) StartTracking(id int, tags []string) {
	_ = ps.pendancyService.StartTracking(id, tags)
}

func (ps *PendancySystem) StopTracking(id int) {
	_ = ps.pendancyService.StopTracking(id)
}

func (ps *PendancySystem) GetCounts(tags []string) int {
	count, _ := ps.pendancyService.GetCounts(tags)
	return count
}
