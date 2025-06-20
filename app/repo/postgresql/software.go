package repository

import (
	"context"

	"github.com/danvei233/softwareMarket-backend/app/domain"
	"gorm.io/gorm"
)

type SQLSoftwareService struct {
	db *gorm.DB
}

func NewSoftwareService(db *gorm.DB) domain.SoftwareService {
	return &SQLSoftwareService{db: db}
}

func (r *SQLSoftwareService) Del(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&domain.Software{}, id).Error
}

func (r *SQLSoftwareService) GetVerList(ctx context.Context, softwareID uint64) ([]domain.Version, error) {
	var vers []domain.Version
	if err := r.db.WithContext(ctx).
		Where("parent_id = ?", softwareID).
		Find(&vers).Error; err != nil {
		return nil, err
	}
	return vers, nil
}

func (r *SQLSoftwareService) Update(ctx context.Context, softwareID uint64, software *domain.Software) error {
	if software.ID == 0 {
		software.ParentID = softwareID
		return r.db.WithContext(ctx).Create(software).Error
	}
	return r.db.WithContext(ctx).
		Model(&domain.Software{}).
		Where("id = ?", software.ID).
		Updates(software).Error
}
