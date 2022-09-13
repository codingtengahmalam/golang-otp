package helper

type UserOtp struct {
	Otp string `json:"otp"`
}

const (
	Numeric = "0123456789"
)

func RandNumeric() string {
	// TODO: Generate Random string
	return "65789"
}
