package getservice

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/danvei233/softwareMarket-backend/app/domain"
	sr "github.com/danvei233/softwareMarket-backend/app/repo/postgresql"
	sqlite "github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// 自动迁移你的模型
	err = db.AutoMigrate(&domain.MainCategory{}, &domain.SubCategory{}, &domain.Software{}, &domain.Version{})
	assert.NoError(t, err)

	return db
}

func TestGetService_GetSoftwareDetail(t *testing.T) {

	db := setupTestDB(t)
	ctx := context.Background()

	// 插入 Main 和 Sub
	main := domain.MainCategory{Name: "Main1"}
	assert.NoError(t, db.Create(&main).Error)
	sub := domain.SubCategory{Name: "Sub1", ParentID: main.ID}
	assert.NoError(t, db.Create(&sub).Error)

	// 插入 Software 和 Version
	sw := domain.Software{
		ParentID:    sub.ID,
		Name:        "TestSoft",
		Description: "desc",
		Author:      "AuthorX",
	}
	assert.NoError(t, db.Create(&sw).Error)

	ver := domain.Version{
		ParentID:      sw.ID,
		VersionNumber: "v1.0",
		Size:          100,
		Action:        1,
		BinaryURL:     "https://cdn.com/bin",
	}
	assert.NoError(t, db.Create(&ver).Error)

	// 注入 repo
	getSvc := &GetService{
		db: db,
		r: sqlRepo{
			sw: sr.NewSoftwareRepo(db), // 实现了 domain.SoftwareRepository
		},
	}

	// 测试 GetSoftwareDetail
	result, err := getSvc.GetSoftwareDetail(ctx, sw.ID)
	assert.NoError(t, err)
	assert.Equal(t, sw.Name, result.Name)
	assert.Len(t, result.Versions, 1)

	data, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("Software Detail:\n%s", data)
	fmt.Println(string(data))
}

func TestGetService_GetAllSoftware(t *testing.T) {
	db := setupTestDB(t)
	ctx := context.Background()

	// 构造测试数据
	main := domain.MainCategory{Name: "MainX"}
	assert.NoError(t, db.Create(&main).Error)
	sub := domain.SubCategory{Name: "SubX", ParentID: main.ID}
	assert.NoError(t, db.Create(&sub).Error)
	sw := domain.Software{Name: "SWX", ParentID: sub.ID}
	assert.NoError(t, db.Create(&sw).Error)

	getSvc := &GetService{
		db: db,
		r: sqlRepo{
			main: sr.NewMainCategoryRepo(db), // 实现了 domain.MainCategoryRepository
		},
	}

	// 调用
	list, err := getSvc.GetAllSoftWareShortcut(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, list)

	data, _ := json.MarshalIndent(list, "", "  ")
	t.Logf("All Software:\n%s", data)
	fmt.Println(string(data))
}

func TestGetService_GetSubList(t *testing.T) {
	db := setupTestDB(t)
	ctx := context.Background()

	// 构造测试数据
	main := domain.MainCategory{Name: "MainCategory"}
	assert.NoError(t, db.Create(&main).Error)

	// 创建多个子类别
	sub1 := domain.SubCategory{Name: "SubCategory1", ParentID: main.ID}
	assert.NoError(t, db.Create(&sub1).Error)
	sub2 := domain.SubCategory{Name: "SubCategory2", ParentID: main.ID}
	assert.NoError(t, db.Create(&sub2).Error)

	// 使用构造函数创建服务
	mainRepo := sr.NewMainCategoryRepo(db)
	getSvc := NewGetService(db, mainRepo, nil, nil, nil)

	// 调用GetSubList方法
	subList, err := getSvc.GetSubList(ctx, main.ID)
	assert.NoError(t, err)
	assert.NotNil(t, subList)
	assert.Len(t, *subList, 2)

	// 验证子类别数据
	categories := *subList
	assert.Contains(t, []string{categories[0].Name, categories[1].Name}, "SubCategory1")
	assert.Contains(t, []string{categories[0].Name, categories[1].Name}, "SubCategory2")

	data, _ := json.MarshalIndent(subList, "", "  ")
	t.Logf("SubCategory List:\n%s", data)
	fmt.Println(string(data))
}
