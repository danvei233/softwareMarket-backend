package domain

import (
	"context"
)

type MainCategory struct {
	ID            uint64        `gorm:"primaryKey;autoIncrement"`
	Name          string        `gorm:"type:varchar(255);not null"`
	SubCategories []SubCategory `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type MainCategoryRepository interface {
	GetBigStrctUntilSoftware(ctx context.Context) ([]MainCategory, error)
	Update(ctx context.Context, o MainCategory) error
	Del(ctx context.Context, id uint64) error
	GetSubList(ctx context.Context, id uint64) ([]SubCategory, error)
}
