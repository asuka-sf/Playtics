package domain

import "errors"

var ErrDuplicateEmail = errors.New("email already exists")
var ErrNotFound = errors.New("resource not found")
var ErrValidation = errors.New("validation error")
