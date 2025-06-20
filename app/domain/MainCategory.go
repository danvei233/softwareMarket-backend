package domain

import "context"

type MainCategory struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(255);not null"`
}
	type MainCategoryRepository interface {
		New(ctx context.Context, o MainCategory) (id uint64, err error)
		GetList(ctx context.Context) ([]MainCategory, error)
		Update(ctx context.Context, o MainCategory) error
		Del(ctx context.Context, id uint64) error
		GetSubList(ctx context.Context, id uint64) ([]SubCategory, error)
		AddSub(ctx context.Context, id uint64, sub SubCategory) (sid uint64, err error)
	}
