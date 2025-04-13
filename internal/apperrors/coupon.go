package apperrors

import "errors"

var (
	ErrUserAlreadyIssued  = errors.New("user already issued coupon")
	ErrCouponCodeConflict = errors.New("coupon code conflict")
)
