package errors

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmailExisted = errors.New("email already exists")
	ErrOTPInvalid   = errors.New("invalid OTP code")
	ErrOTPExpired   = errors.New("OTP code has expired")
)
