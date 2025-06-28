package final

import (
	"github.com/danvei233/softwareMarket-backend/app/handler"
	repository "github.com/danvei233/softwareMarket-backend/app/repo/postgresql"
	getservice "github.com/danvei233/softwareMarket-backend/app/service/Getservice"
	"github.com/danvei233/softwareMarket-backend/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupAPIService(api *gin.RouterGroup, config *utils.AppConfig) error {
	log := utils.GetLog()
	dsn := config.GetDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")

	}
	v := repository.NewVersionRepo(db)
	sw := repository.NewSoftwareRepo(db)
	main := repository.NewMainCategoryRepo(db)
	sub := repository.NewSubCategoryRepo(db)
	service := getservice.NewGetService(db, main, sub, sw, v)
	Handler := handler.NewGetHandeler(service, api.Group("v1"))
	GHandler, err := handler.NewGraphqlHandler(db, api.Group("graphql"), config)
	if err != nil {
		return err
	}
	_ = GHandler
	_ = Handler
	return nil

}
