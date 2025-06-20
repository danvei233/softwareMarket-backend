package domain

import (
	"context"
)

type SubCategory struct {
	Id         uint64
	Name       string
	Icon       string
	ParentId   uint64  
}

type SubCategoryRepository interface {
	Update(ctx context.Context, o *SubCategory) error
	Delete(ctx context.Context, id uint64) error
	GetSwList(ctx context.Context) ([]*SubCategory, error)
	AddSw(ctx context.Context, o *SubCategory) (uint64, error)
}
