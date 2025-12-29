package implements

import (
	"context"
	"fmt"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/models"
	"go-ecommerce-backend-api/internal/repository"
	"go-ecommerce-backend-api/internal/services"
	"go-ecommerce-backend-api/pkg/errors"
	"go-ecommerce-backend-api/pkg/request"
	"go-ecommerce-backend-api/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type authServiceImpl struct {
	userRepo repository.UserRepository
	otpRepo  repository.OTPRepository
}

func NewAuthService(userRepo repository.UserRepository, otpRepo repository.OTPRepository) services.AuthService {
	return &authServiceImpl{
		userRepo: userRepo,
		otpRepo:  otpRepo,
	}
}

func (as *authServiceImpl) GetMe(ctx context.Context, userID string) (*models.User, error) {
	user, err := as.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("Truy xuất thông tin người dùng thất bại: %w", err)
	}

	if user == nil {
		return nil, errors.ErrUserNotFound
	}

	return user, nil
}

func (as *authServiceImpl) Register(ctx context.Context, req request.RegisterRequest) error {
	exists, err := as.userRepo.CheckEmailExists(ctx, req.Email)

	if err != nil {
		return err
	}

	if exists {
		return errors.ErrEmailExisted
	}

	hashedPassword, err := utils.HashAndSalt([]byte(req.Password))

	if err != nil {
		return fmt.Errorf("Hash mật khẩu thất bại: %w", err)
	}

	user := &models.User{
		UserID:   uuid.NewString(),
		Fullname: req.Fullname,
		Email:    req.Email,
		Password: hashedPassword,
		IsActive: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := as.userRepo.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("Tạo account user thất bại: %w", err)
	}

	otpCode, err := utils.GenerateOTP()
	if err != nil {
		return fmt.Errorf("Tạo mã OTP thất bại: %w", err)
	}

	if err := utils.SendOTPEmail(user.Email, user.Fullname, otpCode); err != nil {
		fmt.Printf("Gửi email OTP thất bại: %v\n", err)
	}

	if err := as.otpRepo.SetOTP(ctx, user.UserID, otpCode, 5*time.Minute); err != nil {
		return fmt.Errorf("Lưu mã OTP thất bại: %w", err)
	}

	return nil
}

func (as *authServiceImpl) VerifyOTP(ctx context.Context, req request.VerifyOTPRequest) error {
	savedOTP, err := as.otpRepo.GetOTP(ctx, req.Email)

	if err != nil {
		return errors.ErrOTPExpired
	}

	if savedOTP != req.OTP {
		return errors.ErrOTPInvalid
	}

	if err := as.userRepo.ActiveUserByEmail(ctx, req.Email); err != nil {
		return fmt.Errorf("Kích hoạt tài khoản thất bại: %w", err)
	}

	if err := as.otpRepo.DeleteOTP(ctx, req.Email); err != nil {
		return fmt.Errorf("Xoá mã OTP thất bại: %w", err)
	}

	return nil
}

func (as *authServiceImpl) ResendOTP(ctx context.Context, req request.ResendOTPRequest) error {
	exists, err := as.userRepo.CheckEmailExists(ctx, req.Email)

	if err != nil {
		return err
	}

	if !exists {
		return errors.ErrUserNotFound
	}

	otpCode, err := utils.GenerateOTP()
	if err != nil {
		return fmt.Errorf("Tạo mã OTP thất bại: %w", err)
	}

	fullName, err := as.userRepo.GetNameByEmail(ctx, req.Email)

	if err != nil {
		return fmt.Errorf("Lấy tên người dùng thất bại: %w", err)
	}

	if err := utils.SendOTPEmail(req.Email, fullName, otpCode); err != nil {
		fmt.Printf("Gửi email OTP thất bại: %v\n", err)
	}

	if err := as.otpRepo.SetOTP(ctx, req.Email, otpCode, 5*time.Minute); err != nil {
		return fmt.Errorf("Lưu mã OTP thất bại: %w", err)
	}

	return nil
}

func (as *authServiceImpl) Login(ctx context.Context, req request.LoginRequest) (*models.User, string, string, error) {
	user, err := as.userRepo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return nil, "", "", fmt.Errorf("Truy xuất thông tin người dùng bằng email thất bại")
	}

	if !user.IsActive {
		return nil, "", "", fmt.Errorf("Tài khoản chưa kích hoạt")
	}

	if check := utils.ComparePasswords(user.Password, []byte(req.Password)); !check {
		return nil, "", "", errors.ErrInvalidUser
	}

	accessToken, err := utils.GenerateToken(user.UserID, "customer", global.Config.JWT.AccessExpiry, global.Config.JWT.AccessSecret)
	if err != nil {
		return nil, "", "", fmt.Errorf("Tạo access token thất bại: %w", err)
	}

	refreshToken, err := utils.GenerateToken(user.UserID, "customer", global.Config.JWT.RefreshExpiry, global.Config.JWT.RefreshSecret)
	if err != nil {
		return nil, "", "", fmt.Errorf("Tạo refresh token thất bại: %w", err)
	}
	
	return user, accessToken, refreshToken, nil
}

func (as *authServiceImpl) RefreshToken(ctx context.Context, userID string, userRole string) (string, string, error) {
	accessToken, err := utils.GenerateToken(userID, userRole, global.Config.JWT.AccessExpiry, global.Config.JWT.AccessSecret)
	if err != nil {
		return "", "", fmt.Errorf("Tạo access token thất bại: %w", err)
	}

	refreshToken, err := utils.GenerateToken(userID, userRole, global.Config.JWT.RefreshExpiry, global.Config.JWT.RefreshSecret)
	if err != nil {
		return "", "", fmt.Errorf("Tạo refresh token thất bại: %w", err)
	}

	return accessToken, refreshToken, nil
}

func (as *authServiceImpl) ForgotPassword(ctx context.Context, req request.ForgotPasswordRequest) error {
	exists, err := as.userRepo.CheckEmailExists(ctx, req.Email)

	if err != nil {
		return fmt.Errorf("Kiểm tra người dùng bằng email thất bại: %w", err)
	}

	if !exists {
		return errors.ErrUserNotFound
	}

	userID, err := as.userRepo.GetUserIDByEmail(ctx, req.Email)

	if err != nil {
		return fmt.Errorf("Lấy userID bằng email thất bại: %w", err)
	}

	resetToken, err := utils.GenerateToken(userID, "customer", global.Config.JWT.ResetPasswordExpiry, global.Config.JWT.ResetPasswordSecret)

	if err := as.otpRepo.SetResetPasswordToken(ctx, resetToken, userID); err != nil {
		return fmt.Errorf("Lưu token thất bại: %w", err)
	}

	resetLink := fmt.Sprintf("http://localhost:3000/reset-password?token=%s", resetToken)

	go func() {
		if err := utils.SendResetPasswordEmail(req.Email, resetLink); err != nil {
			fmt.Printf("Gửi email reset password thất bại: %v\n", err)
		}
	}()

	return nil
}

func (as *authServiceImpl) ResetPassword(ctx context.Context, req request.ResetPasswordRequest) error {
	userID, err := as.otpRepo.GetResetPasswordToken(ctx, req.Token)

	if err != nil || userID == "" {
		return errors.ErrTokenInvalid
	}

	hashedPassword, err := utils.HashAndSalt([]byte(req.NewPassword))
	if err != nil {
		return fmt.Errorf("Hash mật khẩu thất bại: %w", err)
	}

	if err := as.userRepo.UpdatePasswordByUserID(ctx, userID, hashedPassword);  err != nil {
		return fmt.Errorf("Cập nhật mật khẩu thất bại: %w", err)
	}

	if err := as.otpRepo.DeleteResetPasswordToken(ctx, req.Token); err != nil {
		return fmt.Errorf("Xoá token thất bại: %w", err)
	}

	return nil
}

func (as *authServiceImpl) ChangePassword(ctx context.Context, userID string, req request.ChangePasswordRequest) error {
	user, err := as.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return fmt.Errorf("Truy xuất thông tin người dùng thất bại: %w", err)
	}

	if user == nil {
		return errors.ErrUserNotFound
	}

	if check := utils.ComparePasswords(user.Password, []byte(req.OldPassword)); !check {
		return errors.ErrInvalidUser
	}

	hashedPassword, err := utils.HashAndSalt([]byte(req.NewPassword))
	if err != nil {
		return fmt.Errorf("Hash mật khẩu thất bại: %w", err)
	}

	if err := as.userRepo.UpdatePasswordByUserID(ctx, userID, hashedPassword);  err != nil {
		return fmt.Errorf("Cập nhật mật khẩu thất bại: %w", err)
	}

	return nil
}