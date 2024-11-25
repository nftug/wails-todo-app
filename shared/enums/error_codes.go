package enums

type ErrorCode string

const (
	ValidationError = ErrorCode("InvalidArg")
	NotFoundError   = ErrorCode("NotFound")
)

var ErrorCodes = []ErrorCode{ValidationError, NotFoundError}

func (e ErrorCode) TSName() string { return string(e) }
