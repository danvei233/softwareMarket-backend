package domain

import (
	"context"
)

type Software struct {
	ParentID    uint64     `gorm:"not null;index"`             // refers to SubCategory.ID
	ID          uint64     `gorm:"primaryKey;autoIncrement"`   // 软件ID
	Name        string     `gorm:"type:varchar(255);not null"` // 软件名称
	Type        uint8      `gorm:"not null"`                   // 软件类型
	Icon        string     `gorm:"type:varchar(255)"`          // 图标
	Description string     `gorm:"type:text"`                  // 描述
	Rate        uint8      `gorm:"not null;default:0"`         // 评分
	DownloadNum uint64     `gorm:"not null;default:0"`         // 下载数量
	Images      []string   `gorm:"type:json;serializer:json"`  // 图片 JSON
	Author      string     `gorm:"type:varchar(100)"`          // 作者
	Document    string     `gorm:"type:varchar(255)"`          // 文档 URL
	CommentURL  string     `gorm:"type:varchar(255)"`          // 评论区 URL
	Meta        []MetaData `gorm:"type:json;serializer:json"`  // 元数据 JSON 存储
}

type MetaData struct {
	Key   string `gorm:"type:varchar(100);not null"`
	Value string `gorm:"type:text"`
}


type SoftwareService interface {
	Del(ctx context.Context, id uint64) error
	GetVerList(ctx context.Context, softwareID uint64) ([]Version, error)
	Update(ctx context.Context, softwareID uint64, software *Software) error
}
