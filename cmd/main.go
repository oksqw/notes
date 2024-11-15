package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"notes"
	"notes/pkg/handler"
	"notes/pkg/initializer"
	"notes/pkg/repository"
	"notes/pkg/service"
)

func main() {
	if err := initializer.InitializeLogger(); err != nil {
		log.Fatal(err)
	}

	if err := initializer.InitConfig(); err != nil {
		logrus.Fatal(err)
	}

	var notesDb, err = initializer.InitializeNotesSQLiteDb(viper.GetString("db.notes.sqlite.filename"))
	if err != nil {
		logrus.Fatal(err)
	}

	notesRepo := repository.NewSQLiteNoteRepository(notesDb)
	repositories := repository.NewRepositories(notesRepo)
	services := service.NewServices(repositories)
	handlers := handler.NewHandler(services)

	server := new(notes.Server)
	if err = server.Run(viper.GetString("port"), handlers.InitializeRouters()); err != nil {
		logrus.Fatal(err)
	}
}
