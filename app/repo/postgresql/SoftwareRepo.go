package repository

import (
	"context"

	"github.com/danvei233/softwareMarket-backend/app/domain"
	"gorm.io/gorm"
)

type SQLSoftwareService struct {
	db dt
}

func NewSoftwareRepo(db *gorm.DB) domain.SoftwareRepository {
	return &SQLSoftwareService{db: dt{db: db}}
}

func (r *SQLSoftwareService) Del(ctx context.Context, id uint64) error {
	return r.db.WithTransaction(ctx).WithContext(ctx).Delete(&domain.Software{}, id).Error
}

func (r *SQLSoftwareService) GetVerList(ctx context.Context, softwareID uint64) (*[]domain.Version, error) {
	var vers []domain.Version
	if err := r.db.WithTransaction(ctx).WithContext(ctx).
		Where("parent_id = ?", softwareID).
		Find(&vers).Error; err != nil {
		return nil, err
	}
	return &vers, nil
}

func (r *SQLSoftwareService) Update(ctx context.Context, softwareID uint64, software *domain.Software) error {
	if software.ID == 0 {
		software.ParentID = softwareID
		return r.db.WithTransaction(ctx).WithContext(ctx).Create(software).Error
	}
	return r.db.WithTransaction(ctx).WithContext(ctx).
		Model(&domain.Software{}).
		Where("id = ?", software.ID).
		Updates(software).Error
}
func (r *SQLSoftwareService) GetSoftwareDetail(ctx context.Context, id uint64) (*domain.Software, error) {
	var software domain.Software
	if err := r.db.WithTransaction(ctx).WithContext(ctx).Preload("Versions").
		First(&software, id).Error; err != nil {
		return nil, err
	}
	return &software, nil
}
