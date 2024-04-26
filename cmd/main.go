package main

import (
	"AiServer/configs"
	"AiServer/internal/models"
	"AiServer/internal/router"
	"AiServer/pkg/log"
	"AiServer/pkg/utils"
)

func main() {
	log.InitLog()
	err := configs.InitCfg()
	if err != nil {
		panic(err)
	}
	err = utils.InitDbEngine(configs.DatabaseConfig)
	if err != nil {
		panic(err)
	}
	models.Migrate()

	utils.InitRedisClient(configs.RedisConfig)

	r := router.NewRouter()
	r.Run(":8080")
}
