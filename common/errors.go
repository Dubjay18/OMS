package common

import "errors"

var (
	ErrorNoItems = errors.New("items must be at least one")
)
