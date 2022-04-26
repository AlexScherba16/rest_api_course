package main

import (
	"fmt"
	"net/http"
	"rest_api_course/config"
	"rest_api_course/internal/comment"
	"rest_api_course/internal/database"
	transportHTTP "rest_api_course/internal/transport/http"
)

// App - the struct which contains things like
// pointers to database connections
type App struct{}

// Run - handles the startup of our application
func (app *App) Run() error {

	appConfig := config.NewAppConfig()

	var err error
	db, err := database.NewDatabase(&appConfig.DBConfig)
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up")
		fmt.Println(err)
	}
}
