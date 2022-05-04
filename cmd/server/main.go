package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"rest_api_course/config"
	"rest_api_course/internal/comment"
	"rest_api_course/internal/database"
	transportHTTP "rest_api_course/internal/transport/http"
)

// App - contains application info
type App struct {
	Name    string
	Version string
}

// Run - handles the startup of our application
func (app *App) Run() error {

	appConfig := config.NewAppConfig()
	//log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up application")

	var err error
	db, err := database.NewDatabase(&appConfig.DBConfig)
	if err != nil {
		log.Error(err)
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		log.Error(err)
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	app := App{
		Name:    "Comments REST service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error Starting Up", err)
	}
}
