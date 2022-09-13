package request

type (
	CreateUserRequest struct {
		Name     string `json:"name"  validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	GetOtpRequest struct {
		Email string `json:"email" validate:"required"`
	}

	LoginWithOTPRequest struct {
		Email string `json:"email" validate:"required"`
		Otp   string `json:"otp" validate:"required"`
	}
)
