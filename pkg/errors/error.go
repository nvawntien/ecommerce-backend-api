package errors

import "errors"

var (
	ErrInvalidUser = errors.New("Email hoặc mật khẩu không đúng")
	ErrUserNotFound = errors.New("Không tìm thấy người dùng")
	ErrTokenInvalid = errors.New("Token không hợp lệ")
	ErrEmailExisted = errors.New("Email đã tồn tại")
	ErrOTPInvalid   = errors.New("Mã OTP không hợp lệ")
	ErrOTPExpired   = errors.New("Mã OTP đã hết hạn")
)
