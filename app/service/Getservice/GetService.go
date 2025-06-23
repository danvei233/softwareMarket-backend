package getservice

import (
	"context"

	"github.com/danvei233/softwareMarket-backend/app/domain"
	"gorm.io/gorm"
)

type sqlRepo struct {
	main domain.MainCategoryRepository
	sub  domain.SubCategoryRepository
	sw   domain.SoftwareRepository
	v    domain.VersionRepository
}
type GetService struct {
	r  sqlRepo
	db *gorm.DB
}

func (s *GetService) GetAllSoftWare(ctx context.Context) (*[]domain.MainCategory, error) {
	var mc *[]domain.MainCategory
	var err error

	err = s.db.Transaction(func(tx *gorm.DB) error {
		ctxv := context.WithValue(ctx, "tx", tx)
		mc, err = s.r.main.GetBigStrctUntilSoftware(ctxv)
		return err
	})
	if err != nil {
		return nil, err
	}
	return mc, err
}
func (s *GetService) GetSoftwareDetail(ctx context.Context, id uint64) (*domain.Software, error) {
	var sw *domain.Software
	var err error
	err = s.db.Transaction(func(tx *gorm.DB) error {

		ctxv := context.WithValue(ctx, "tx", tx)
		sw, err = s.r.sw.GetSoftwareDetail(ctxv, id)
		return err
	})
	if err != nil {
		return nil, err
	}
	return sw, err
}

// GetAllSoftWare
// GetSoftwareDetail
