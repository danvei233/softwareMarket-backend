package repository

import (
	"context"

	"github.com/Swmarket/app/domain"
	"gorm.io/gorm"
)

type SQLMainCategoryRepo struct {
	db *gorm.DB
}

// NewMainCategoryRepo creates a MainCategoryRepository backed by GORM.
func NewMainCategoryRepo(db *gorm.DB) domain.MainCategoryRepository {
	return &SQLMainCategoryRepo{db: db}
}

// GetBigStrctUntilSoftware returns all main categories with their subcategories and softwares (no versions).
func (r *SQLMainCategoryRepo) GetBigStrctUntilSoftware(ctx context.Context) ([]domain.MainCategory, error) {
	var mains []domain.MainCategory
	if err := r.db.WithContext(ctx).
		Preload("SubCategories").
		Find(&mains).Error; err != nil {
		return nil, err
	}
	// Load softwares for each subcategory
	for i := range mains {
		for j := range mains[i].SubCategories {
			var list []*domain.Software
			if err := r.db.WithContext(ctx).
				Where("parent_id = ?", mains[i].SubCategories[j].ID).
				Find(&list).Error; err != nil {
				return nil, err
			}
			mains[i].SubCategories[j].Softwares = list
		}
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
