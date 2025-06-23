package repository

import (
	"context"

	"github.com/danvei233/softwareMarket-backend/app/domain"
	"gorm.io/gorm"
)

type SQLSubCategoryRepo struct {
	db dt
}

func NewSubCategoryRepo(db *gorm.DB) domain.SubCategoryRepository {
	return &SQLSubCategoryRepo{db: dt{db: db}}
}

func (r *SQLSubCategoryRepo) Update(ctx context.Context, o *domain.SubCategory) error {
	if o.ID == 0 {
		return r.db.WithTransaction(ctx).WithContext(ctx).Create(o).Error
	}

	return r.db.WithTransaction(ctx).WithContext(ctx).
		Model(&domain.SubCategory{}).
		Where("id = ?", o.ID).
		Updates(o).Error
}

func (r *SQLSubCategoryRepo) Del(ctx context.Context, id uint64) error {
	return r.db.WithTransaction(ctx).WithContext(ctx).Delete(&domain.SubCategory{}, id).Error
}

func (r *SQLSubCategoryRepo) GetSoftwareList(ctx context.Context, subCategoryID uint64) ([]*domain.Software, error) {
	var list []*domain.Software
	if err := r.db.WithTransaction(ctx).WithContext(ctx).
		Where("parent_id = ?", subCategoryID).
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
