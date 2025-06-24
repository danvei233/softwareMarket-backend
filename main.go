package main

import (
	"github.com/danvei233/softwareMarket-backend/app/handler"
	repository "github.com/danvei233/softwareMarket-backend/app/repo/postgresql"
	getservice "github.com/danvei233/softwareMarket-backend/app/service/Getservice"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	router := gin.Default()

	// Configure router settings
	api := router.Group("api/v1/")

	Handler, err := SetupAPIService(api)
	if err != nil {
		log.Fatal("Failed to setup service: ", err)
	}

	// Keep handler in scope
	_ = Handler

	// Start HTTP server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
func SetupAPIService(api *gin.RouterGroup) (*handler.GetHandler, error) {
	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	v := repository.NewVersionRepo(db)
	sw := repository.NewSoftwareRepo(db)
	main := repository.NewMainCategoryRepo(db)
	sub := repository.NewSubCategoryRepo(db)
	service := getservice.NewGetService(db, main, sub, sw, v)
	Handler := handler.NewGetHandeler(service, api)
	return Handler, nil

}
