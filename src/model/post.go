package model

import (
	"context"
	"golang-otp/src/request"
	"golang-otp/src/response"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) (*Post, error)
	UpdateByID(ctx context.Context, id int, post *Post) (*Post, error)
	FindByID(ctx context.Context, id int) (*Post, error)
	Delete(ctx context.Context, id int) error
	Fetch(ctx context.Context, limit, offset int) ([]*Post, int64, error)
}

type PostUsecase interface {
	GetPostList(ctx context.Context, limit, offset int) (*response.PostsResponse, error)
	GetPostByID(ctx context.Context, id int) (*response.PostResponse, error)
	StorePost(ctx context.Context, request request.PostRequest) (*response.PostResponse, error)
}
