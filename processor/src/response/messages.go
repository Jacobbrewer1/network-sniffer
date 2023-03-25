package response

import "errors"

var (
	InvalidApiKey      = errors.New("invalid API-Key")
	Unauthorised       = errors.New("unauthorised")
	InvalidContentType = errors.New("invalid Content-Type")
	NotFound           = errors.New("not found")
	MethodNotAllowed   = errors.New("method not allowed")
)
