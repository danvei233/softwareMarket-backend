package handler

//一坨屎
//我不会
//postman
//QAQ

//import (
//	"context"
//	"encoding/json"
//	"errors"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//	"unsafe"   //这傻逼ai怎么几把调用unsafe了
//被角力
//
//	"github.com/danvei233/softwareMarket-backend/app/domain"
//	getservice "github.com/danvei233/softwareMarket-backend/app/service/Getservice"
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//)
//
//// MockMainCategoryRepo mocks the MainCategoryRepository
//type MockMainCategoryRepo struct {
//	getBigStrctUntilSoftwareFn func(ctx context.Context) (*[]domain.MainCategory, error)
//	getSubListFn               func(ctx context.Context, id uint64) (*[]domain.SubCategory, error)
//	updateFn                   func(ctx context.Context, o domain.MainCategory) error
//	delFn                      func(ctx context.Context, id uint64) error
//}
//
//func (m *MockMainCategoryRepo) RetrieveMainCategoryDetails(ctx context.Context) (*[]domain.MainCategory, error) {
//	return m.getBigStrctUntilSoftwareFn(ctx)
//}
//
//func (m *MockMainCategoryRepo) GetSubListByMainId(ctx context.Context, id uint64) (*[]domain.SubCategory, error) {
//	return m.getSubListFn(ctx, id)
//}
//
//func (m *MockMainCategoryRepo) Update(ctx context.Context, o domain.MainCategory) error {
//	if m.updateFn != nil {
//		return m.updateFn(ctx, o)
//	}
//	return nil
//}
//
//func (m *MockMainCategoryRepo) Del(ctx context.Context, id uint64) error {
//	if m.delFn != nil {
//		return m.delFn(ctx, id)
//	}
//	return nil
//}
//
//// MockSubCategoryRepo mocks the SubCategoryRepository
//type MockSubCategoryRepo struct {
//	updateFn          func(ctx context.Context, o *domain.SubCategory) error
//	delFn             func(ctx context.Context, id uint64) error
//	getSoftwareListFn func(ctx context.Context, subCategoryID uint64) ([]*domain.Software, error)
//}
//
//func (m *MockSubCategoryRepo) Update(ctx context.Context, o *domain.SubCategory) error {
//	if m.updateFn != nil {
//		return m.updateFn(ctx, o)
//	}
//	return nil
//}
//
//func (m *MockSubCategoryRepo) Del(ctx context.Context, id uint64) error {
//	if m.delFn != nil {
//		return m.delFn(ctx, id)
//	}
//	return nil
//}
//
//func (m *MockSubCategoryRepo) GetSoftwareList(ctx context.Context, subCategoryID uint64) ([]*domain.Software, error) {
//	if m.getSoftwareListFn != nil {
//		return m.getSoftwareListFn(ctx, subCategoryID)
//	}
//	return nil, nil
//}
//
//// MockSoftwareRepo mocks the SoftwareRepository
//type MockSoftwareRepo struct {
//	getSoftwareDetailFn func(ctx context.Context, id uint64) (*domain.Software, error)
//	delFn               func(ctx context.Context, id uint64) error
//	getVerListFn        func(ctx context.Context, softwareID uint64) (*[]domain.Version, error)
//	updateFn            func(ctx context.Context, softwareID uint64, software *domain.Software) error
//}
//
//func (m *MockSoftwareRepo) GetSoftwareDetail(ctx context.Context, id uint64) (*domain.Software, error) {
//	return m.getSoftwareDetailFn(ctx, id)
//}
//
//func (m *MockSoftwareRepo) Del(ctx context.Context, id uint64) error {
//	if m.delFn != nil {
//		return m.delFn(ctx, id)
//	}
//	return nil
//}
//
//func (m *MockSoftwareRepo) GetVerList(ctx context.Context, softwareID uint64) (*[]domain.Version, error) {
//	if m.getVerListFn != nil {
//		return m.getVerListFn(ctx, softwareID)
//	}
//	return nil, nil
//}
//
//func (m *MockSoftwareRepo) Update(ctx context.Context, softwareID uint64, software *domain.Software) error {
//	if m.updateFn != nil {
//		return m.updateFn(ctx, softwareID, software)
//	}
//	return nil
//}
//
//// MockVersionRepo mocks the VersionRepository
//type MockVersionRepo struct {
//	updateFn func(ctx context.Context, v *domain.Version) error
//	delFn    func(ctx context.Context, id uint64) error
//}
//
//func (m *MockVersionRepo) Update(ctx context.Context, v *domain.Version) error {
//	if m.updateFn != nil {
//		return m.updateFn(ctx, v)
//	}
//	return nil
//}
//
//func (m *MockVersionRepo) Del(ctx context.Context, id uint64) error {
//	if m.delFn != nil {
//		return m.delFn(ctx, id)
//	}
//	return nil
//}
//
//// setupTest initializes a test environment with a mock service and gin router
//func setupTest() (*gin.Engine, *gin.RouterGroup) {
//	gin.SetMode(gin.TestMode)
//	router := gin.Default()
//	group := router.Group("/api")
//	return router, group
//}
//
//// MockGetService is a mock implementation of the GetService
//type MockGetService struct {
//	mainRepo domain.MainCategoryRepository
//	swRepo   domain.SoftwareRepository
//}
//
//// GetAllSoftWareShortcut mocks the GetAllSoftWareShortcut method without using DB transaction
//func (m *MockGetService) GetAllSoftWareShortcut(ctx context.Context) (*[]domain.MainCategory, error) {
//	if m.mainRepo == nil {
//		return nil, errors.New("main repository is nil")
//	}
//	return m.mainRepo.RetrieveMainCategoryDetails(ctx)
//}
//
//// GetSoftwareDetail mocks the GetSoftwareDetail method without using DB transaction
//func (m *MockGetService) GetSoftwareDetail(ctx context.Context, id uint64) (*domain.Software, error) {
//	if m.swRepo == nil {
//		return nil, errors.New("software repository is nil")
//	}
//	return m.swRepo.GetSoftwareDetail(ctx, id)
//}
//
//// GetSubListByMainId mocks the GetSubListByMainId method without using DB transaction
//func (m *MockGetService) GetSubListByMainId(ctx context.Context, id uint64) (*[]domain.SubCategory, error) {
//	if m.mainRepo == nil {
//		return nil, errors.New("main repository is nil")
//	}
//	return m.mainRepo.GetSubListByMainId(ctx, id)
//}
//
//// createMockService creates a MockGetService with mock repositories
//func createMockService(
//	mainRepo domain.MainCategoryRepository,
//	subRepo domain.SubCategoryRepository,
//	swRepo domain.SoftwareRepository,
//	vRepo domain.VersionRepository,
//) *getservice.GetService {
//	// Cast our MockGetService to *getservice.GetService
//	// This is unsafe but works for testing since we control the handler's usage
//	return (*getservice.GetService)(unsafe.Pointer(&MockGetService{
//		mainRepo: mainRepo,
//		swRepo:   swRepo,
//	}))
//}
//
//// TestGetAllSoftware tests the GetAllSoftware handler
//func TestGetAllSoftware(t *testing.T) {
//	router, group := setupTest()
//
//	t.Run("Success", func(t *testing.T) {
//		// Prepare mock data
//		mainCategories := &[]domain.MainCategory{
//			{
//				ID:   1,
//				Name: "Category 1",
//				SubCategories: []domain.SubCategory{
//					{
//						ID:   1,
//						Name: "SubCategory 1",
//					},
//				},
//			},
//		}
//
//		// Create mock repositories
//		mockMainRepo := &MockMainCategoryRepo{
//			getBigStrctUntilSoftwareFn: func(ctx context.Context) (*[]domain.MainCategory, error) {
//				return mainCategories, nil
//			},
//		}
//
//		// Create service with mock repositories
//		service := createMockService(mockMainRepo, nil, nil, nil)
//
//		// Initialize handler with service
//		NewGetHandeler(service, group)
//
//		// Create test request
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("GET", "/api/getlist", nil)
//		router.ServeHTTP(w, req)
//
//		// Assert response
//		assert.Equal(t, http.StatusOK, w.Code)
//
//		var response map[string]interface{}
//		err := json.Unmarshal(w.Body.Bytes(), &response)
//		assert.NoError(t, err)
//		assert.Equal(t, "ok", response["msg"])
//		assert.NotNil(t, response["list"])
//	})
//
//	t.Run("Error", func(t *testing.T) {
//		// Reset router and group for this test
//		router, group := setupTest()
//
//		// Create mock repositories with error
//		mockMainRepo := &MockMainCategoryRepo{
//			getBigStrctUntilSoftwareFn: func(ctx context.Context) (*[]domain.MainCategory, error) {
//				return nil, errors.New("database error")
//			},
//		}
//
//		// Create service with mock repositories
//		service := createMockService(mockMainRepo, nil, nil, nil)
//
//		// Initialize handler with service
//		NewGetHandeler(service, group)
//
//		// Create test request
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("GET", "/api/getlist", nil)
//		router.ServeHTTP(w, req)
//
//		// Assert response
//		assert.Equal(t, http.StatusInternalServerError, w.Code)
//
//		var response map[string]interface{}
//		err := json.Unmarshal(w.Body.Bytes(), &response)
//		assert.NoError(t, err)
//		assert.Contains(t, response, "msg")
//	})
//}
//
//// TestGetSoftware tests the GetSoftware handler
//func TestGetSoftware(t *testing.T) {
//	t.Run("Success", func(t *testing.T) {
//		// Setup router for this test
//		router, group := setupTest()
//
//		// Prepare mock data
//		software := &domain.Software{
//			ID:          1,
//			Name:        "Test Software",
//			Description: "Test Description",
//			Author:      "Test Author",
//			Versions: []domain.Version{
//				{
//					ID:            1,
//					VersionNumber: "1.0.0",
//				},
//			},
//		}
//
//		// Create mock repositories
//		mockSoftwareRepo := &MockSoftwareRepo{
//			getSoftwareDetailFn: func(ctx context.Context, id uint64) (*domain.Software, error) {
//				if id == 1 {
//					return software, nil
//				}
//				return nil, errors.New("not found")
//			},
//		}
//
//		// Create service with mock repositories
//		service := createMockService(nil, nil, mockSoftwareRepo, nil)
//
//		// Initialize handler with service
//		NewGetHandeler(service, group)
//
//		// Create test request
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("GET", "/api/getsoftare?id=1", nil)
//		router.ServeHTTP(w, req)
//
//		// Assert response
//		assert.Equal(t, http.StatusOK, w.Code)
//
//		var response map[string]interface{}
//		err := json.Unmarshal(w.Body.Bytes(), &response)
//		assert.NoError(t, err)
//		assert.Equal(t, "ok", response["msg"])
//		assert.NotNil(t, response["software"])
//	})
//
//	t.Run("Invalid ID", func(t *testing.T) {
//		// Setup router for this test
//		router, group := setupTest()
//
//		// Create a minimal service with empty repositories
//		// The handler will fail before calling any repository methods
//		service := createMockService(nil, nil, nil, nil)
//
//		// Initialize handler with service
//		NewGetHandeler(service, group)
//
//		// Create test request with invalid ID
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("GET", "/api/getsoftare?id=invalid", nil)
//		router.ServeHTTP(w, req)
//
//		// Assert response
//		assert.Equal(t, http.StatusInternalServerError, w.Code)
//
//		var response map[string]interface{}
//		err := json.Unmarshal(w.Body.Bytes(), &response)
//		assert.NoError(t, err)
//		assert.Contains(t, response, "msg")
//		assert.Contains(t, response["msg"], "No valid id in query")
//	})
//
//	t.Run("Service Error", func(t *testing.T) {
//		// Setup router for this test
//		router, group := setupTest()
//
//		// Create mock repositories with error
//		mockSoftwareRepo := &MockSoftwareRepo{
//			getSoftwareDetailFn: func(ctx context.Context, id uint64) (*domain.Software, error) {
//				return nil, errors.New("database error")
//			},
//		}
//
//		// Create service with mock repositories
//		service := createMockService(nil, nil, mockSoftwareRepo, nil)
//
//		// Initialize handler with service
//		NewGetHandeler(service, group)
//
//		// Create test request
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("GET", "/api/getsoftare?id=2", nil)
//		router.ServeHTTP(w, req)
//
//		// Assert response
//		assert.Equal(t, http.StatusInternalServerError, w.Code)
//
//		var response map[string]interface{}
//		err := json.Unmarshal(w.Body.Bytes(), &response)
//		assert.NoError(t, err)
//		assert.Contains(t, response, "msg")
//	})
//}
//
//// TestGetSubCategory tests the GetSubcategoryList handler
//// Note: In the actual implementation, the GetSoftware handler is registered for the "getsubcategory" route
//func TestGetSubCategory(t *testing.T) {
//	t.Run("Success", func(t *testing.T) {
//		// Setup router for this test
//		router, group := setupTest()
//
//		// Prepare mock data
//		software := &domain.Software{
//			ID:          1,
//			Name:        "Test Software",
//			Description: "Test Description",
//			Author:      "Test Author",
//		}
//
//		// Create mock repositories
//		mockSoftwareRepo := &MockSoftwareRepo{
//			getSoftwareDetailFn: func(ctx context.Context, id uint64) (*domain.Software, error) {
//				if id == 1 {
//					return software, nil
//				}
//				return nil, errors.New("not found")
//			},
//		}
//
//		// Create service with mock repositories
//		service := createMockService(nil, nil, mockSoftwareRepo, nil)
//
//		// Initialize handler with service
//		NewGetHandeler(service, group)
//
//		// Create test request
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("GET", "/api/getsubcategory?id=1", nil)
//		router.ServeHTTP(w, req)
//
//		// Assert response
//		assert.Equal(t, http.StatusOK, w.Code)
//
//		var response map[string]interface{}
//		err := json.Unmarshal(w.Body.Bytes(), &response)
//		assert.NoError(t, err)
//		assert.Equal(t, "ok", response["msg"])
//		assert.NotNil(t, response["software"]) // Note: The response key is "software", not "subcategory"
//	})
//
//	t.Run("Invalid ID", func(t *testing.T) {
//		// Setup router for this test
//		router, group := setupTest()
//
//		// Create a minimal service with empty repositories
//		// The handler will fail before calling any repository methods
//		service := createMockService(nil, nil, nil, nil)
//
//		// Initialize handler with service
//		NewGetHandeler(service, group)
//
//		// Create test request with invalid ID
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("GET", "/api/getsubcategory?id=invalid", nil)
//		router.ServeHTTP(w, req)
//
//		// Assert response
//		assert.Equal(t, http.StatusInternalServerError, w.Code)
//
//		var response map[string]interface{}
//		err := json.Unmarshal(w.Body.Bytes(), &response)
//		assert.NoError(t, err)
//		assert.Contains(t, response, "msg")
//		assert.Contains(t, response["msg"], "No valid id in query")
//	})
//
//	t.Run("Service Error", func(t *testing.T) {
//		// Setup router for this test
//		router, group := setupTest()
//
//		// Create mock repositories with error
//		mockSoftwareRepo := &MockSoftwareRepo{
//			getSoftwareDetailFn: func(ctx context.Context, id uint64) (*domain.Software, error) {
//				return nil, errors.New("database error")
//			},
//		}
//
//		// Create service with mock repositories
//		service := createMockService(nil, nil, mockSoftwareRepo, nil)
//
//		// Initialize handler with service
//		NewGetHandeler(service, group)
//
//		// Create test request
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("GET", "/api/getsubcategory?id=2", nil)
//		router.ServeHTTP(w, req)
//
//		// Assert response
//		assert.Equal(t, http.StatusInternalServerError, w.Code)
//
//		var response map[string]interface{}
//		err := json.Unmarshal(w.Body.Bytes(), &response)
//		assert.NoError(t, err)
//		assert.Contains(t, response, "msg")
//	})
//}
