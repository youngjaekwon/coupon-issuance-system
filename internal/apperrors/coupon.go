package apperrors

import "errors"

var (
	ErrUserAlreadyIssued = errors.New("user already issued coupon")
)
