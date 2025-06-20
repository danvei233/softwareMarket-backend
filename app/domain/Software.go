package domain

import (
	"context"
)

type Software struct {
	ParentId    uint64
	Id          uint64   // 软件ID
	Name        string   // 软件名称
	Type        uint8    // 软件类型
	Icon        string   // 图标
	Description string   // 描述
	Rate        uint8    // 评分
	DownloadNum uint64   // 下载数量
	Images      []string // 图片
	Author      string   // 作者
	Document    string   // 文档
	CommentURL  string   // 评论区url
	Meta        MetaData
}

type MetaData struct {
	Key   string
	Value string
}

// SoftwareService defines the interface for software operations with context.
type SoftwareService interface {
	// Del deletes a software by its ID.
	Del(ctx context.Context, id uint64) error

	// GetVerList returns a list of versions for the software.
	GetVerList(ctx context.Context, softwareID uint64) ([]Version, error)

	// AddVer adds a new version to the software and returns its ID.
	AddVer(ctx context.Context, softwareID uint64, version Version) (uint64, error)
}
