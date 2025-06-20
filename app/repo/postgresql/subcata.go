package repository

import (
	"context"

	"github.com/Swmarket/app/domain"
	"gorm.io/gorm"
)

type SQLSubCategoryRepo struct {
	db *gorm.DB
}

func NewSubCategoryRepo(db *gorm.DB) domain.SubCategoryRepository {
	return &SQLSubCategoryRepo{db: db}
}

func (r *SQLSubCategoryRepo) Update(ctx context.Context, o *domain.SubCategory) error {
	if o.ID == 0 {
		return r.db.WithContext(ctx).Create(o).Error
	}



	return r.db.WithContext(ctx).
		Model(&domain.SubCategory{}).
		Where("id = ?", o.ID).
		Updates(o).Error
}

func (r *SQLSubCategoryRepo) Del(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&domain.SubCategory{}, id).Error
}

func (r *SQLSubCategoryRepo) GetSoftwareList(ctx context.Context, subCategoryID uint64) ([]*domain.Software, error) {
	var list []*domain.Software
	if err := r.db.WithContext(ctx).
		Where("parent_id = ?", subCategoryID).
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
