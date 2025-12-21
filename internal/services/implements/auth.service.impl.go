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

	accessToken, err := utils.GenerateToken(user.UserID, "access", global.Config.JWT.AccessExpiry, global.Config.JWT.AccessSecret)
	if err != nil {
		return nil, "", "", fmt.Errorf("Tạo access token thất bại: %w", err)
	}

	refreshToken, err := utils.GenerateToken(user.UserID, "refresh", global.Config.JWT.RefreshExpiry, global.Config.JWT.RefreshSecret)
	if err != nil {
		return nil, "", "", fmt.Errorf("Tạo refresh token thất bại: %w", err)
	}
	
	return user, accessToken, refreshToken, nil
}
