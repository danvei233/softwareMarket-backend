package final

import (
	"github.com/danvei233/softwareMarket-backend/app/handler"
	repository "github.com/danvei233/softwareMarket-backend/app/repo/postgresql"
	"github.com/danvei233/softwareMarket-backend/app/service/DownloadSerivce"
	getservice "github.com/danvei233/softwareMarket-backend/app/service/Getservice"
	"github.com/danvei233/softwareMarket-backend/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupAPIService(api *gin.RouterGroup, config *utils.AppConfig) error {
	log := utils.GetLog()
	db, err := gorm.Open(postgres.Open(config.GetDsn()), &gorm.Config{})

	sw := repository.NewSoftwareRepo(db)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")

	}
	_ = handler.NewGetHandeler(
		getservice.NewGetService(db,
			repository.NewMainCategoryRepo(db),
			repository.NewSubCategoryRepo(db),
			sw,
			repository.NewVersionRepo(db)),
		api.Group("api/v1"))
	_, err = handler.NewGraphqlHandler(db, api.Group("graphql"), config)
	if err != nil {
		return err
	}
	_, err = handler.NewDownloadSerivce(api.Group("public"),
		DownloadSerivce.NewDownloadService(db, sw))
	if err != nil {
		return err
	}
	return nil

}
