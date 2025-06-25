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

// NewGetService creates a new instance of GetService with the given repositories and database connection
func NewGetService(db *gorm.DB, main domain.MainCategoryRepository, sub domain.SubCategoryRepository, sw domain.SoftwareRepository, v domain.VersionRepository) *GetService {
	return &GetService{
		db: db,
		r: sqlRepo{
			main: main,
			sub:  sub,
			sw:   sw,
			v:    v,
		},
	}
}
func (s *GetService) GetMainCategory(ctx context.Context) (*[]domain.MainCategory, error) {
	var mains *[]domain.MainCategory
	var err error
	err = s.db.Transaction(func(tx *gorm.DB) error {
		ctxv := context.WithValue(ctx, "tx", tx)
		mains, err = s.r.main.GetMainCategoryList(ctxv)
		return err
	})
	if err != nil {
		return nil, err
	}
	return mains, err
}

func (s *GetService) GetAllSoftWareShortcut(ctx context.Context, id uint64, subPage, subLimit int, softPage, softLimit int) (*domain.MainCategory, error) {
	var mc *domain.MainCategory
	var err error

	err = s.db.Transaction(func(tx *gorm.DB) error {
		ctxv := context.WithValue(ctx, "tx", tx)
		mc, err = s.r.main.RetrieveMainCategoryDetails(ctxv, id, subPage, subLimit, softPage, softLimit)
		return err
	})
	if err != nil {
		return nil, err
	}
	return mc, err
}

func (s *GetService) GetSoftwareFromSubcategory(ctx context.Context, id uint64, sublimt, subpage int) (*domain.SubCategory, error) {
	var sub *domain.SubCategory
	var err error
	err = s.db.Transaction(func(tx *gorm.DB) error {
		ctxv := context.WithValue(ctx, "tx", tx)
		sub, err = s.r.sub.GetSoftwareList(ctxv, id, subpage, sublimt)
		return err
	})
	if err != nil {
		return nil, err
	}
	return sub, nil
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

// GetAllSoftWareShortcut
// GetSoftwareDetail
// GetSubListByMainId

func (s *GetService) GetSubList(ctx context.Context, id uint64) (*[]domain.SubCategory, error) {
	var subList *[]domain.SubCategory
	var err error

	err = s.db.Transaction(func(tx *gorm.DB) error {
ctxV := context.WithValue(ctx, "tx", tx)
		subList, err = s.r.main.GetSubListByMainId(ctxv, id)
		return err
	})
	if err != nil {
		return nil, err
	}
	return subList, err
}
