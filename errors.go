package main

import "errors"

var (
	FailToAuthError      = errors.New("fail to auth")
	TokenExpireError     = errors.New("token expire")
	EventIsNotStatsError = errors.New("event is not status")
	ArgsIsEmptyError     = errors.New("args is empty")
)
