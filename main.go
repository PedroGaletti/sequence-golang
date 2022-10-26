package main

import (
	"challenge/cmd/sequence"
	"challenge/configs"
	"challenge/db"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	env = configs.Env
)

func main() {
	// * Configure *
	log.Info().Msg("Setting logs configure...")
	zerolog.TimeFieldFormat = "02/01/2006 15:04:05"
	logLevel, _ := zerolog.ParseLevel(env.LogLevel)
	zerolog.SetGlobalLevel(logLevel)

	// * Database *
	log.Info().Msg("Configuration databases...")
	database, err := db.InitMyqlDb(env.SqlUser, env.SqlPassword, env.SqlHost, env.SqlPort, env.SqlDb)
	if err != nil {
		log.Panic().Msg(fmt.Sprintf("Database connection error: %s", err))
	}

	log.Info().Msg("Running migrations...")
	db.Migrate(database)

	// * Routes and Dependencies *
	log.Info().Msg("Creating routes and dependecies...")
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
	}))

	sequence.InjectDependency(router.Group("/sequence"), database)

	router.Run()
}
