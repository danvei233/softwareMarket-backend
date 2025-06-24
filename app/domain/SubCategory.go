package domain

import (
	"context"
)

type SubCategory struct {
	ID        uint64      `gorm:"primaryKey;autoIncrement"`
	Name      string      `gorm:"type:varchar(255);not null"`
	Icon      string      `gorm:"type:varchar(255)"`
	ParentID  uint64      `gorm:"not null;index"`
	Softwares []*Software `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type SubCategoryRepository interface {
	Update(ctx context.Context, o *SubCategory) error
	Del(ctx context.Context, id uint64) error
	GetSoftwareList(ctx context.Context, subCategoryID uint64, softPage int, softLimit int) (*SubCategory, error)
}
