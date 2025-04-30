package constant

const (
	InternalServerErrorMessage     = "system busy, please try again later"
	EOFErrorMessage                = "missing body request"
	JsonSyntaxErrorMessage         = "invalid JSON syntax"
	JsonUnMarshallTypeErrorMessage = "invalid value for %s"
	ValidationErrorMessage         = "input validation error"
	RequestTimeoutErrorMessage     = "failed to process request in time, please try again"
	RequestCanceledErrorMessage    = "request canceled by client, please try again"
)
