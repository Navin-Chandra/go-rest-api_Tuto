package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/Navin-Chandra/go-rest-api_Tuto/internal/transport/http"
)

// App - the struct which contains things like
// pointers to database connections
type App struct{}

// Run - handles the startup of our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	// fmt.Println("Go REST API Course")\
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up")
		fmt.Println(err)
	}
}
