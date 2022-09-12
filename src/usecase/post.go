package usecase

import (
	"context"
	"golang-otp/src/model"
	"golang-otp/src/request"
	"golang-otp/src/response"
	"strings"
)

type postUsecase struct {
	postRepository model.PostRepository
}

func NewPostUsecase(post model.PostRepository) model.PostUsecase {
	return &postUsecase{postRepository: post}
}

func (p *postUsecase) GetPostList(ctx context.Context, limit, offset int) (*response.PostsResponse, error) {
	posts, total, err := p.postRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	resp := new(response.PostsResponse)

	resp.Post = posts
	resp.Meta.Offset = offset
	resp.Meta.Limit = limit
	resp.Meta.Total = total

	return resp, nil
}

func (p *postUsecase) GetPostByID(ctx context.Context, id int) (*response.PostResponse, error) {
	post, err := p.postRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := new(response.PostResponse)
	resp.Post = post

	return resp, nil

}

func (p *postUsecase) StorePost(ctx context.Context, request request.PostRequest) (*response.PostResponse, error) {
	postSlug := strings.Replace(request.Title, " ", "-", 0)

	newPost := model.Post{
		Title:   request.Title,
		Slug:    strings.ToLower(postSlug),
		Content: request.Content,
	}

	post, err := p.postRepository.Create(ctx, &newPost)
	if err != nil {
		return nil, err
	}

	resp := new(response.PostResponse)
	resp.Post = post

	return resp, nil
}
