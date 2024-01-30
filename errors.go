package main

import "errors"

var (
	FailToAuthError  = errors.New("fail to auth")
	TokenExpireError = errors.New("token expire")
)
