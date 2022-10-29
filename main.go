package main

import (
	"konecta/config"
	"konecta/feature"
	"konecta/repository"
	"konecta/server"
	"log"
)

func run() error {
	cfg, err := config.NewConfig()

	if err != nil {
		return err
	}

	db, err := config.InitDB(cfg)

	if err != nil {
		return err
	}

	repository := repository.NewRepository(db.Database("konecta"))
	feature := feature.NewFeature(repository)

	srv := server.NewServer(cfg, feature)

	return srv.Run()
}

func main() {
	err := run()

	if err != nil {
		log.Panic(err)
	}
}
