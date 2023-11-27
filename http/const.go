package http

type methodName string

const (
	MethodGet     methodName = "GET"
	MethodHead    methodName = "HEAD"
	MethodPost    methodName = "POST"
	MethodPut     methodName = "PUT"
	MethodPatch   methodName = "PATCH" // RFC 5789
	MethodDelete  methodName = "DELETE"
	MethodConnect methodName = "CONNECT"
	MethodOptions methodName = "OPTIONS"
	MethodTrace   methodName = "TRACE"
)
