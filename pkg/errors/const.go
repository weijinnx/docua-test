package errors

import (
	errs "errors"
)

var (
	// EmptyMessageError var
	EmptyMessageError = errs.New("empty message")
	// TimeParseError var
	TimeParseError = errs.New("fail to parse time from query params")
	// EmptyTimeError var
	EmptyTimeError = errs.New("empty time")
	// FailConnectRedisError var
	FailConnectRedisError = errs.New("fail connect to redis")
)
