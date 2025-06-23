package repository

import (
	"context"

	"github.com/danvei233/softwareMarket-backend/app/domain"
	"gorm.io/gorm"
)

type SQLMainCategoryRepo struct {
	db dt
}
type dt struct {
	db *gorm.DB
}

func NewMainCategoryRepo(db *gorm.DB) domain.MainCategoryRepository {
	return &SQLMainCategoryRepo{db: dt{db: db}}
}

func (d *dt) WithTransaction(ctx context.Context) *gorm.DB {
	tx := ctx.Value("tx")
	if txDB, ok := tx.(*gorm.DB); ok && txDB != nil {

		return txDB

	}
	return d.db
}

func (r *SQLMainCategoryRepo) GetBigStrctUntilSoftware(ctx context.Context) (*[]domain.MainCategory, error) {
	var mains []domain.MainCategory
	if err := r.db.WithTransaction(ctx).WithContext(ctx).
		Preload("SubCategories.Softwares").
		Find(&mains).Error; err != nil {
		return nil, err
	}
	return &mains, nil
}

func (r *SQLMainCategoryRepo) Update(ctx context.Context, o domain.MainCategory) error {
	if o.ID == 0 {
		return r.db.WithTransaction(ctx).WithContext(ctx).Create(&o).Error
	}
	return r.db.WithTransaction(ctx).WithContext(ctx).
		Model(&domain.MainCategory{}).
		Where("id = ?", o.ID).
		Update("name", o.Name).Error
}

func (r *SQLMainCategoryRepo) Del(ctx context.Context, id uint64) error {
	return r.db.WithTransaction(ctx).WithContext(ctx).Delete(&domain.MainCategory{}, id).Error
}

func (r *SQLMainCategoryRepo) GetSubList(ctx context.Context, id uint64) (*[]domain.SubCategory, error) {
	var subs []domain.SubCategory

	if err := r.db.WithTransaction(ctx).WithContext(ctx).
		Where("parent_id = ?", id).
		Find(&subs).Error; err != nil {
		return nil, err
	}
	return &subs, nil
}
