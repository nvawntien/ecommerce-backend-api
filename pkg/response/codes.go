package response

const (
	//1xxx - Auth
	CodeSuccess          = 1000 // Thành công
	CodeInternalError    = 1001 // Lỗi hệ thống không xác định
	CodeInvalidParams    = 1002 // Tham số không hợp lệ
	CodeUnauthorized     = 1003 // Chưa đăng nhập hoặc token hết hạn
	CodePermissionDenied = 1004 // Không có quyền truy cập
	CodeTooManyRequests  = 1005 // Quá nhiều yêu cầu, vui lòng thử lại sau

	// 2xxx - User & Account
	CodeEmailExisted    = 2001 // Email đã tồn tại
	CodeUsernameExisted = 2002 // Tên đăng nhập đã tồn tại
	CodeUserNotFound    = 2003 // Người dùng không tồn tại
	CodeWrongPassword   = 2004 // Mật khẩu không đúng
	CodeOTPInvalid      = 2005 // Mã OTP không hợp lệ
	CodeOTPExpired      = 2006 // Mã OTP hết hạn
)

var msg = map[int]string{
	CodeSuccess:           "Thành công",
	CodeInternalError:    "Lỗi hệ thống không xác định",
	CodeInvalidParams:    "Tham số không hợp lệ",
	CodeUnauthorized:     "Chưa đăng nhập hoặc token hết hạn",
	CodePermissionDenied: "Không có quyền truy cập",
	CodeTooManyRequests:  "Quá nhiều yêu cầu, vui lòng thử lại sau",

	CodeEmailExisted:    "Email đã tồn tại",
	CodeUsernameExisted: "Tên đăng nhập đã tồn tại",
	CodeUserNotFound:    "Người dùng không tồn tại",
	CodeWrongPassword:   "Mật khẩu không đúng",
}
