package errors

import "net/http"

var (
	ErrNotFound         = New(http.StatusNotFound, "notfound", "not found", nil)
	ErrResourceNotFound = New(http.StatusNotFound, "resource.notfound", "resource not found", nil)
	ErrLoginFailure     = New(http.StatusUnauthorized, "login.failure", "username and password do not match", nil)
)
