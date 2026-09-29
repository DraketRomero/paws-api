package domain

import "net/http"

type StatusCode int

const (
	SUCCESS      StatusCode = http.StatusOK
	CREATED      StatusCode = http.StatusCreated
	NOT_FOUND    StatusCode = http.StatusNotFound
	BADREQUEST   StatusCode = http.StatusBadRequest
	FORBBIDEN    StatusCode = http.StatusForbidden
	UNAUTHORIZED StatusCode = http.StatusUnauthorized
)
