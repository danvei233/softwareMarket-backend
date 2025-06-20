package domain

import "context"

type Version struct {
	ParentId      uint64
	Id            uint64
	VersionNumber string
	Size          uint64
	Action        uint16
	BinaryURL     string
}

type VersionRepository interface {
	Update(ctx context.Context, v *Version) error
	Del(ctx context.Context, id uint64) error
	Get(ctx context.Context, id uint64) (*Version, error)
}
