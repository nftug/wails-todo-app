package enums

type ErrorCode string

const (
	InvalidArgError = ErrorCode("InvalidArg")
	NotFoundError   = ErrorCode("NotFound")
)

var ErrorCodes = []ErrorCode{InvalidArgError, NotFoundError}

func (e ErrorCode) TSName() string { return string(e) }
