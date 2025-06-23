package repository

import (
	"context"

	"github.com/danvei233/softwareMarket-backend/app/domain"
	"gorm.io/gorm"
)

type SQLVersionRepo struct {
	db dt
}

func NewVersionRepo(db *gorm.DB) domain.VersionRepository {
	return &SQLVersionRepo{db: dt{db: db}}
}

func (r *SQLVersionRepo) Update(ctx context.Context, v *domain.Version) error {
	if v.ID == 0 {
		return r.db.WithTransaction(ctx).WithContext(ctx).Create(v).Error
	}
	return r.db.WithTransaction(ctx).WithContext(ctx).
		Model(&domain.Version{}).
		Where("id = ?", v.ID).
		Updates(v).Error
}

func (r *SQLVersionRepo) Del(ctx context.Context, id uint64) error {
	return r.db.WithTransaction(ctx).WithContext(ctx).Delete(&domain.Version{}, id).Error
}
