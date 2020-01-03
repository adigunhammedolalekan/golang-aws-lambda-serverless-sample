package repos

import (
	"errors"
	"github.com/adigunhammedolalekan/golang-aws-lambda-sample/types"
	"github.com/jinzhu/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db:db}
}

func (repo *PostRepository) CreatePost(p *types.Post) (*types.Post, error) {
	if p.Title == "" || p.Body == "" || p.User == "" {
		return nil, errors.New("invalid post, post parameter is empty")
	}
	if err := repo.db.Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (repo *PostRepository) GetPost(id uint) (*types.Post, error) {
	p := &types.Post{}
	err := repo.db.Table("posts").Where("id = ?", id).First(p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (repo *PostRepository) EditPost(id uint, p *types.Post) (*types.Post, error) {
	if err := repo.db.Table("posts").Where("id = ?", id).Update(p).Error; err != nil {
		return nil, err
	}
	return repo.GetPost(id)
}

func (repo *PostRepository) DeletePost(id uint) error {
	return repo.db.Table("posts").Where("id = ?", id).Delete(&types.Post{}).Error
}
