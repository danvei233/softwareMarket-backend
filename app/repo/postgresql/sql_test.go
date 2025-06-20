// repository/sql_test.go
package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	sqlite "github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/danvei233/softwareMarket-backend/app/domain"
)

// setupTestDB creates an in-memory SQLite database and migrates domain models.
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
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
	assert.Len(t, result, 100)

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
	for i := 1; i <= 20; i++ {
		mc := domain.MainCategory{Name: fmt.Sprintf("BMMain-%d", i)}
		db.Create(&mc)
		for j := 1; j <= 5; j++ {
			sc := domain.SubCategory{ParentID: mc.ID, Name: fmt.Sprintf("BMSub-%d-%d", i, j)}
			db.Create(&sc)
			for k := 1; k <= 5; k++ {
				db.Create(&domain.Software{ParentID: sc.ID, Name: fmt.Sprintf("BMSoft-%d-%d-%d", i, j, k)})
			}
		}
	}
	repo := NewMainCategoryRepo(db)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		repo.GetBigStrctUntilSoftware(ctx)
	}
}
