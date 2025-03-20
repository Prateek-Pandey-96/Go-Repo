package facade

type ConcreteFacade struct {
	// interfaces will be used in actual prod ready code
	OrderSubsystem   *OrderSubsystem
	CookingSubsystem *CookingSubsystem
	*ServingSubsystem
}

func GetRestaurantFacade(orderSubsystem *OrderSubsystem,
	cookingSubsystem *CookingSubsystem,
	servingSubsystem *ServingSubsystem) IFacade {

	return &ConcreteFacade{
		OrderSubsystem:   orderSubsystem,
		CookingSubsystem: cookingSubsystem,
		ServingSubsystem: servingSubsystem,
	}
}

func (cf *ConcreteFacade) GetFood() {
	cf.OrderSubsystem.TakeOrder()
	cf.CookingSubsystem.CookFood()
	cf.ServingSubsystem.ServeOrder()
}
