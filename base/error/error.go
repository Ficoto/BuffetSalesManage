package err_enum

import "errors"

// Custom error code
var (
	ErrInvalidArgs     = errors.New("invalid argument")
	ErrInvalidObjectID = errors.New("invalid object id")
)
