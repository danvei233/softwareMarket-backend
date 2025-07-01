package DownloadSerivce

import (
	"context"
	"github.com/danvei233/softwareMarket-backend/app/domain"
	"gorm.io/gorm"
)

type DownloadService struct {
	r  *gorm.DB
	sw domain.SoftwareRepository
}

func NewDownloadService(r *gorm.DB, sw domain.SoftwareRepository) *DownloadService {
	return &DownloadService{r: r, sw: sw}
}
func (s *DownloadService) AddDownloadNum(ctx context.Context, id uint64) error {
	err := s.r.Transaction(func(tx *gorm.DB) error {
		context.WithValue(ctx, "tx", tx)
		item, err := s.sw.GetSoftwareDetail(ctx, id)
		if err != nil {
			return err
		}
		item.DownloadNum += 1
		err = s.sw.Update(ctx, id, item)
		if err != nil {
			return err
		}
		return nil

	})
	return err
}
