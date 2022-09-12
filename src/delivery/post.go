package delivery

import (
	"github.com/labstack/echo/v4"
	"golang-otp/src/model"
	"golang-otp/src/request"
	"net/http"
	"strconv"
)

type postDelivery struct {
	postUsecase model.PostUsecase
}

type PostDelivery interface {
	Mount(group *echo.Group)
}

func NewPostDelivery(postUsecase model.PostUsecase) PostDelivery {
	return &postDelivery{postUsecase: postUsecase}
}

func (p *postDelivery) Mount(group *echo.Group) {
	group.GET("", p.FetchPostHandler)
	group.POST("", p.StorePostHandler)
	group.GET("/:id", p.DetailPostHandler)
	group.DELETE("/:id", p.DeletePostHandler)
	group.PATCH("/:id", p.EditPostHandler)
}

func (p *postDelivery) FetchPostHandler(c echo.Context) error {
	ctx := c.Request().Context()

	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	postList, err := p.postUsecase.GetPostList(ctx, limitInt, offsetInt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, postList)

}

func (p *postDelivery) StorePostHandler(c echo.Context) error {
	var req request.PostRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := c.Validate(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, true)

}

func (p *postDelivery) DetailPostHandler(c echo.Context) error {
	panic("implement this")

}

func (p *postDelivery) DeletePostHandler(c echo.Context) error {
	panic("implement this")

}

func (p *postDelivery) EditPostHandler(c echo.Context) error {
	panic("implement this")
}
