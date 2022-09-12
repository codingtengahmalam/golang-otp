package repository

import (
	"context"
	"encoding/json"
	"golang-otp/config"
	"golang-otp/src/model"
	"strconv"
)

type postRepository struct {
	Cfg config.Config
}

func NewPostRepository(cfg config.Config) model.PostRepository {
	return &postRepository{Cfg: cfg}
}

func (p *postRepository) FindByID(ctx context.Context, id int) (*model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postRepository) Create(ctx context.Context, post *model.Post) (*model.Post, error) {
	if err := p.Cfg.Database().WithContext(ctx).Create(&post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (p *postRepository) UpdateByID(ctx context.Context, id int, post *model.Post) (*model.Post, error) {
	if err := p.Cfg.Database().WithContext(ctx).
		Model(&model.Post{ID: id}).Updates(post).Find(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (p *postRepository) Delete(ctx context.Context, id int) error {
	err := p.Cfg.Database().WithContext(ctx).Delete(&model.Post{ID: id}).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *postRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.Post, int64, error) {
	var data []*model.Post
	var count int64
	p.Cfg.Database().WithContext(ctx).Model(&model.Post{}).Count(&count)

	key := "article:limit" + strconv.Itoa(limit) + ":offset:" + strconv.Itoa(offset)
	posts, err := p.Cfg.Redis().Get(ctx, key)
	if err != nil {
		if err := p.Cfg.Database().WithContext(ctx).
			Select("id", "title", "slug", "content", "created_at", "updated_at").
			Limit(limit).Offset(offset).Find(&data).Error; err != nil {
			return nil, 0, err
		}

		err = p.Cfg.Redis().Set(ctx, key, data)
		if err != nil {
			return nil, 0, err
		}

		return data, count, nil
	}

	err = json.Unmarshal([]byte(posts), &data)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}
