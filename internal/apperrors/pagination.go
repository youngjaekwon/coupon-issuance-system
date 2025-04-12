package apperrors

import "errors"

var ErrInvalidPage = errors.New("invalid page number")
var ErrInvalidPageSize = errors.New("invalid page size")
