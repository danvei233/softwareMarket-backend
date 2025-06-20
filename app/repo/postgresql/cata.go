package repository

import (
	"context"

	"github.com/danvei233/softwareMarket-backend/app/domain"
	"gorm.io/gorm"
)

type SQLMainCategoryRepo struct {
	db *gorm.DB
}

func NewMainCategoryRepo(db *gorm.DB) domain.MainCategoryRepository {
	return &SQLMainCategoryRepo{db: db}
}

func (r *SQLMainCategoryRepo) GetBigStrctUntilSoftware(ctx context.Context) ([]domain.MainCategory, error) {
	var mains []domain.MainCategory
	if err := r.db.WithContext(ctx).
		Preload("SubCategories.Softwares").
		Find(&mains).Error; err != nil {
		return nil, err
	}
	return mains, nil
}

func (r *SQLMainCategoryRepo) Update(ctx context.Context, o domain.MainCategory) error {
	if o.ID == 0 {
		return r.db.WithContext(ctx).Create(&o).Error
	}
	return r.db.WithContext(ctx).
		Model(&domain.MainCategory{}).
		Where("id = ?", o.ID).
		Update("name", o.Name).Error
}

func (r *SQLMainCategoryRepo) Del(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&domain.MainCategory{}, id).Error
}

func (r *SQLMainCategoryRepo) GetSubList(ctx context.Context, id uint64) ([]domain.SubCategory, error) {
	var subs []domain.SubCategory
	if err := r.db.WithContext(ctx).
		Where("parent_id = ?", id).
		Find(&subs).Error; err != nil {
		return nil, err
	}
	return subs, nil
}
