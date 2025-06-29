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
	GetMainCategoryList(ctx context.Context) (*[]MainCategory, error)
	RetrieveMainCategoryDetails(ctx context.Context,
		id uint64,
		subPage, subLimit int,
		softPage, softLimit int, t ...any) (*MainCategory, error)

	Update(ctx context.Context, o MainCategory) error
	Del(ctx context.Context, id uint64) error
	GetSubListByMainId(ctx context.Context, id uint64) (*[]SubCategory, error)
	//GetBigStructUntilSoftware(ctx context.Context) (*[]MainCategory, error)
}
