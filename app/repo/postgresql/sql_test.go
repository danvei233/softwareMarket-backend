// repository/sql_test.go
package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/danvei233/softwareMarket-backend/app/domain"
)

// setupTestDB creates an in-memory SQLite database and migrates domain models.
func setupTestDB(t *testing.T) *gorm.DB {

	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(
		&domain.MainCategory{},
		&domain.SubCategory{},
		&domain.Software{},
		&domain.Version{},
	)
	assert.NoError(t, err)
	return db
}

// TestGetBigStructUntilSoftware_JSON seeds many records and outputs the result as JSON.
func TestGetBigStructUntilSoftware_JSON(t *testing.T) {
	db := setupTestDB(t)
	ctx := context.Background()

	// Seed: 100 main categories, each with 10 subcategories, each with 5 software entries
	for i := 1; i <= 100; i++ {
		mc := domain.MainCategory{Name: fmt.Sprintf("MainCategory-%d", i)}
		assert.NoError(t, db.WithContext(ctx).Create(&mc).Error)
		for j := 1; j <= 10; j++ {
			sc := domain.SubCategory{ParentID: mc.ID, Name: fmt.Sprintf("SubCategory-%d-%d", i, j)}
			assert.NoError(t, db.WithContext(ctx).Create(&sc).Error)
			for k := 1; k <= 5; k++ {
				sw := domain.Software{ParentID: sc.ID, Name: fmt.Sprintf("Software-%d-%d-%d", i, j, k)}
				assert.NoError(t, db.WithContext(ctx).Create(&sw).Error)
			}
		}
	}

	repo := NewMainCategoryRepo(db)
	result, err := repo.GetBigStrctUntilSoftware(ctx)
	assert.NoError(t, err)
	fmt.Print(len(*result))
	assert.Len(t, *result, 100)

	// Marshal to JSON
	data, err := json.Marshal(result)
	assert.NoError(t, err)

	// Output JSON
	t.Logf("GetBigStrctUntilSoftware JSON: %s", data)
	fmt.Println(string(data))
}

// TestGetSoftwareList_JSON seeds a subcategory with many software and outputs JSON.
func TestGetSoftwareList_JSON(t *testing.T) {
	db := setupTestDB(t)
	ctx := context.Background()

	// Seed one main and one subcategory
	mc := domain.MainCategory{Name: "MC"}
	assert.NoError(t, db.WithContext(ctx).Create(&mc).Error)
	sc := domain.SubCategory{ParentID: mc.ID, Name: "SubTest"}
	assert.NoError(t, db.WithContext(ctx).Create(&sc).Error)

	// Seed 50 software entries under this subcategory
	for k := 1; k <= 50; k++ {
		sw := domain.Software{ParentID: sc.ID, Name: fmt.Sprintf("BenchSoftware-%d", k)}
		assert.NoError(t, db.WithContext(ctx).Create(&sw).Error)
	}

	subRepo := NewSubCategoryRepo(db)
	list, err := subRepo.GetSoftwareList(ctx, sc.ID)
	assert.NoError(t, err)
	assert.Len(t, list, 50)

	// Marshal to JSON and output
	data, err := json.Marshal(list)
	assert.NoError(t, err)
	t.Logf("GetSoftwareList JSON: %s", data)
	fmt.Println(string(data))
}

// BenchmarkGetBigStructUntilSoftware runs the get method on large data.
func BenchmarkGetBigStructUntilSoftware(b *testing.B) {
	db := setupTestDB(nil)
	ctx := context.Background()

	// Pre-seed moderate data
	//for i := 1; i <= 15; i++ {
	//	mc := domain.MainCategory{Name: fmt.Sprintf("BMMain-%d", i)}
	//	db.Create(&mc)
	//	for j := 1; j <= 30; j++ {
	//		sc := domain.SubCategory{ParentID: mc.ID, Name: fmt.Sprintf("BMSub-%d-%d", i, j)}
	//		db.Create(&sc)
	//		for k := 1; k <= 500; k++ {
	//			db.Create(&domain.Software{ParentID: sc.ID, Name: fmt.Sprintf("BMSoft-%d-%d-%d", i, j, k)})
	//		}
	//	}
	//}
	repo := NewMainCategoryRepo(db)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		repo.GetBigStrctUntilSoftware(ctx)
	}
	var sw domain.Software
	err := db.WithContext(ctx).First(&sw).Error
	if err == nil {
		b.Logf("Random Software: %+v", sw)
	}
}

// 手动查询实现
func BenchmarkGetBigStructUntilSoftware_RawSQL(b *testing.B) {
	db := setupTestDB(nil)
	ctx := context.Background()

	// // Pre-seed moderate data
	// for i := 1; i <= 15; i++ {
	// 	mc := domain.MainCategory{Name: fmt.Sprintf("BMMain-%d", i)}
	// 	db.Create(&mc)
	// 	for j := 1; j <= 30; j++ {
	// 		sc := domain.SubCategory{ParentID: mc.ID, Name: fmt.Sprintf("BMSub-%d-%d", i, j)}
	// 		db.Create(&sc)
	// 		for k := 1; k <= 500; k++ {
	// 			db.Create(&domain.Software{ParentID: sc.ID, Name: fmt.Sprintf("BMSoft-%d-%d-%d", i, j, k)})
	// 		}
	// 	}
	// }
	type Row struct {
		MainCategoryID   uint
		MainCategoryName string
		SubCategoryID    uint
		SubCategoryName  string
		SoftwareID       uint
		SoftwareName     string
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var rows []Row
		db.WithContext(ctx).Raw(`
			SELECT 
				m.id as main_category_id, m.name as main_category_name,
				s.id as sub_category_id, s.name as sub_category_name,
				sw.id as software_id, sw.name as software_name
			FROM main_categories m
			JOIN sub_categories s ON s.parent_id = m.id
			JOIN software sw ON sw.parent_id = s.id
		`).Scan(&rows)
		// 可选：构建嵌套结构（略）
	}
	var sw domain.Software
	err := db.WithContext(ctx).First(&sw).Error
	if err == nil {
		b.Logf("Random Software: %+v", sw)
	}
}
func BenchmarkGetBigStructUntilSoftware_Manual(b *testing.B) {
	db := setupTestDB(nil)
	ctx := context.Background()

	// Pre-seed moderate data
	// for i := 1; i <= 15; i++ {
	// 	mc := domain.MainCategory{Name: fmt.Sprintf("BMMain-%d", i)}
	// 	db.Create(&mc)
	// 	for j := 1; j <= 30; j++ {
	// 		sc := domain.SubCategory{ParentID: mc.ID, Name: fmt.Sprintf("BMSub-%d-%d", i, j)}
	// 		db.Create(&sc)
	// 		for k := 1; k <= 500; k++ {
	// 			db.Create(&domain.Software{ParentID: sc.ID, Name: fmt.Sprintf("BMSoft-%d-%d-%d", i, j, k)})
	// 		}
	// 	}
	// }
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var mainCategories []domain.MainCategory
		db.WithContext(ctx).Find(&mainCategories)
		for i := range mainCategories {
			var subCategories []domain.SubCategory
			db.WithContext(ctx).Where("parent_id = ?", mainCategories[i].ID).Find(&subCategories)
			mainCategories[i].SubCategories = subCategories
			for j := range subCategories {
				var softwares []*domain.Software
				db.WithContext(ctx).Where("parent_id = ?", subCategories[j].ID).Find(&softwares)
				mainCategories[i].SubCategories[j].Softwares = softwares
			}
		}
	}
	var sw domain.Software
	err := db.WithContext(ctx).First(&sw).Error
	if err == nil {
		b.Logf("Random Software: %+v", sw)
	}
}

// TestGetSoftwareDetail_JSON seeds a software with detailed info and versions, then outputs JSON.
func TestGetSoftwareDetail_JSON(t *testing.T) {
	db := setupTestDB(t)
	ctx := context.Background()

	// 1. Seed main category and subcategory
	mc := domain.MainCategory{Name: "DetailMC"}
	assert.NoError(t, db.WithContext(ctx).Create(&mc).Error)

	sc := domain.SubCategory{
		ParentID: mc.ID,
		Name:     "DetailSub",
	}
	assert.NoError(t, db.WithContext(ctx).Create(&sc).Error)

	// 2. Seed software with images and metadata
	sw := domain.Software{
		ParentID:    sc.ID,
		Name:        "DetailSoftware",
		Type:        1,
		Icon:        "icon.png",
		Description: "A detailed software for testing",
		Rate:        5,
		DownloadNum: 42,
		Images:      []string{"https://cdn.example.com/img1.png", "https://cdn.example.com/img2.png"},
		Author:      "Test Author",
		Document:    "https://example.com/doc",
		CommentURL:  "https://example.com/comment",
		Meta: []domain.MetaData{
			{Key: "OS", Value: "Windows"},
			{Key: "Lang", Value: "Go"},
		},
	}
	assert.NoError(t, db.WithContext(ctx).Create(&sw).Error)

	// 3. Seed two versions
	versions := []domain.Version{
		{
			ParentID:      sw.ID,
			VersionNumber: "1.0.0",
			Size:          1024,
			Action:        1,
			BinaryURL:     "https://example.com/bin/v1.0.0",
		},
		{
			ParentID:      sw.ID,
			VersionNumber: "1.1.0",
			Size:          2048,
			Action:        2,
			BinaryURL:     "https://example.com/bin/v1.1.0",
		},
	}
	for _, v := range versions {
		assert.NoError(t, db.WithContext(ctx).Create(&v).Error)
	}

	// 4. Call repository and verify
	repo := NewSoftwareRepo(db)
	detail, err := repo.GetSoftwareDetail(ctx, sw.ID)
	assert.NoError(t, err)

	// 基本字段验证
	assert.Equal(t, sw.Name, detail.Name)
	assert.Equal(t, sw.Description, detail.Description)
	assert.Equal(t, sw.Author, detail.Author)

	// 元数据和图片
	assert.Len(t, detail.Images, 2)
	assert.Len(t, detail.Meta, 2)
	assert.Contains(t, detail.Meta, domain.MetaData{Key: "OS", Value: "Windows"})

	// 版本列表验证
	assert.Len(t, detail.Versions, 2)
	assert.Equal(t, versions[0].VersionNumber, detail.Versions[0].VersionNumber)
	assert.Equal(t, versions[1].BinaryURL, detail.Versions[1].BinaryURL)

	// 5. Marshal to JSON and print
	data, err := json.MarshalIndent(detail, "", "  ")
	assert.NoError(t, err)
	t.Logf("GetSoftwareDetail JSON: %s", data)
	fmt.Println(string(data))
}
