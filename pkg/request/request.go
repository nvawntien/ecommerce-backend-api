package request

type RegisterRequest struct {
	Fullname string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=6"`
}

type ResendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token           string `json:"token" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
}

type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password" binding:"required,min=8"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
}

type CreateCategoryRequest struct {
	Name     string `json:"name" binding:"required"`
	ParentID *int   `json:"parent_id,omitempty"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	ParentID *int `json:"parent_id,omitempty"`
}

type CreateProductRequest struct {
	CategoryID  int     `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Brand       string  `json:"brand" binding:"required"`
	BasePrice   float64 `json:"base_price" binding:"required,gt=0"`
	Variants	[]struct {
		SKU 		 string  `json:"sku" binding:"required"`
		VariantName   string  `json:"variant_name" binding:"required"`
		PriceModifier float64 `json:"price_modifier"`
		StockQuantity int     `json:"stock_quantity" binding:"required,gte=0"`
		ImageURL      string  `json:"image_url" binding:"required,url"`
	} `json:"variants" binding:"required,dive,required"`
}