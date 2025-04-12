package models

func AllModels() []interface{} {
	return []interface{}{
		&Campaign{},
		&Coupon{},
	}
}
