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
	db, err := gorm.Open(postgres.Open(config.GetDsn()), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")

	}
	_ = handler.NewGetHandeler(
		getservice.NewGetService(db,
			repository.NewMainCategoryRepo(db),
			repository.NewSubCategoryRepo(db),
			repository.NewSoftwareRepo(db),
			repository.NewVersionRepo(db)),
		api.Group("v1"))
	_, err = handler.NewGraphqlHandler(db, api.Group("graphql"), config)
	if err != nil {
		return err
	}
	return nil

}
