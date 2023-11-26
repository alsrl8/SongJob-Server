package web

type HttpResponseCode int

// Success status codes
const (
	StatusOK                   HttpResponseCode = 200
	StatusCreated              HttpResponseCode = 201
	StatusAccepted             HttpResponseCode = 202
	StatusNonAuthoritativeInfo HttpResponseCode = 203
	StatusNoContent            HttpResponseCode = 204
	StatusResetContent         HttpResponseCode = 205
	StatusPartialContent       HttpResponseCode = 206
)

// Redirection status codes
const (
	StatusMultipleChoices  HttpResponseCode = 300
	StatusMovedPermanently HttpResponseCode = 301
	StatusFound            HttpResponseCode = 302
)

// Client error status codes
const (
	StatusBadRequest      HttpResponseCode = 400
	StatusUnauthorized    HttpResponseCode = 401
	StatusPaymentRequired HttpResponseCode = 402
	StatusForbidden       HttpResponseCode = 403
	StatusNotFound        HttpResponseCode = 404
)

// Server error status codes
const (
	StatusInternalServerError HttpResponseCode = 500
	StatusNotImplemented      HttpResponseCode = 501
	StatusBadGateway          HttpResponseCode = 502
)
