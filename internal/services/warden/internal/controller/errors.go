package controller

import "errors"

var (
	ErrMissingRequiredField = errors.New("Missing required fields")
	ErrAccessDenied         = errors.New("Access Denied")
)
