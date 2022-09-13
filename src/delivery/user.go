package delivery

import (
	"github.com/labstack/echo/v4"
	"golang-otp/src/model"
	"golang-otp/src/request"
	"net/http"
)

type userDelivery struct {
	userUsecase model.UserUsecase
}

type UserDelivery interface {
	Mount(group *echo.Group)
}

func NewUserDelivery(UserUsecase model.UserUsecase) UserDelivery {
	return &userDelivery{userUsecase: UserUsecase}
}

func (p *userDelivery) Mount(group *echo.Group) {
	group.POST("/register", p.registerHandler)
	group.POST("/request-otp", p.requestOtpHandler)
	group.POST("/login", p.loginHandler)
}

func (p *userDelivery) requestOtpHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.GetOtpRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := c.Validate(req)
	if err != nil {
		return err
	}

	err = p.userUsecase.RequestOtp(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
	})

}

func (p *userDelivery) loginHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.LoginWithOTPRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	otp, err := p.userUsecase.LoginWithOTP(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   otp,
	})

}

func (p *userDelivery) registerHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	UserList, err := p.userUsecase.CreateUser(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, UserList)
}
