package main

import (
	"github.com/danvei233/softwareMarket-backend/app/final"
	"github.com/danvei233/softwareMarket-backend/app/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	log := utils.GetLog()
	router := gin.Default()
	log.Info().Msg("Server is starting...")

	config, err := utils.NewAppConfig("config/app.ini")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config file.")
		panic(err)
	}
	log.Info().Msg("Config loaded.")

	// Configure router settings
	api := router.Group("api")
	log.Info().Msg("API router registered.")
	err = final.SetupAPIService(api, config)
	log.Info().Msg("API service setup.")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to setup service.")
	}
	// Load App Config

	// Start HTTP server
	if err := router.Run(config.App.Addr + ":" + config.App.Port); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server. ")
	}
}
