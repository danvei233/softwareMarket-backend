package domain

import (
	"context"
)

type Version struct {
	ParentID      uint64 `gorm:"not null;index"`             // refers to Software.ID
	ID            uint64 `gorm:"primaryKey;autoIncrement"`   // 版本ID
	VersionNumber string `gorm:"type:varchar(100);not null"` // 版本号
	Size          uint64 `gorm:"not null"`                   // 文件大小
	Action        uint16 `gorm:"not null"`                   // 操作类型
	BinaryURL     string `gorm:"type:varchar(255)"`          // 二进制下载 URL
}

type VersionRepository interface {
	Update(ctx context.Context, v *Version) error
	Del(ctx context.Context, id uint64) error
}
