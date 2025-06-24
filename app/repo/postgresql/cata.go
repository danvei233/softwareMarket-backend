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
if txDB, ok := ctx.Value("tx").(*gorm.DB); ok || txDB != nil {
	return txDB
}
}
func (r *SQLMainCategoryRepo) GetMainCategoryList(ctx context.Context) (*[]domain.MainCategory, error) {
	var mains []domain.MainCategory
	if err := r.db.WithTransaction(ctx).WithContext(ctx).
		Find(&mains).Error; err != nil {
		return nil, err
	}
	return &mains, nil
}
func (r *SQLMainCategoryRepo) RetrieveMainCategoryDetails(
	ctx context.Context,
	id uint64,
	subPage, subLimit int,
	softPage, softLimit int,
) (*domain.MainCategory, error) {
	var main domain.MainCategory

	err := r.db.WithTransaction(ctx).WithContext(ctx).
		// 只查询这一条主分类
		Where("id = ?", id).
		// 子分类分页
		Preload("SubCategories", func(db *gorm.DB) *gorm.DB {
			return db.
				Limit(subLimit).
				Offset((subPage - 1) * subLimit).
				Order("id ASC") // 按需排序
		}).
		// 每个子分类对应的软件分页
		Preload("SubCategories.Softwares", func(db *gorm.DB) *gorm.DB {
			return db.
				Limit(softLimit).
				Offset((softPage - 1) * softLimit).
				Order("id ASC") // 按需排序
		}).
		First(&main).Error
	if err != nil {
		return nil, err
	}
	return &main, nil
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

func (r *SQLMainCategoryRepo) GetSubListByMainId(ctx context.Context, id uint64) (*[]domain.SubCategory, error) {
	var subs []domain.SubCategory

	if err := r.db.WithTransaction(ctx).WithContext(ctx).
		Where("parent_id = ?", id).
		Find(&subs).Error; err != nil {
		return nil, err
	}
	return &subs, nil
}

// GetBigStructUntilSoftware retrieves all main categories with their subcategories and software
func (r *SQLMainCategoryRepo) GetBigStructUntilSoftware(ctx context.Context) (*[]domain.MainCategory, error) {
	var mainCategories []domain.MainCategory

	if err := r.db.WithTransaction(ctx).WithContext(ctx).
		Preload("SubCategories").
		Preload("SubCategories.Softwares").
		Find(&mainCategories).Error;err != nil {
		return nil, err
	}
	return &mainCategories, nil
}
