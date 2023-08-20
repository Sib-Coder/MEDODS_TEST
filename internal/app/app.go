package app

import (
	"awesomeProject13/internal/endpoint"
	"awesomeProject13/internal/service"
	"awesomeProject13/internal/storage"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	database *storage.Database
	endpoint *endpoint.Endpoint
	service  *service.Service
	echo     *echo.Echo
}

func New() (*App, error) {
	app := &App{}
	app.database.Db = storage.New()
	app.service = service.New(app.database)
	app.endpoint = endpoint.New(app.service)
	app.echo = echo.New()
	//endpoit
	app.echo.GET("/jwt", app.endpoint.GetJwt)
	//app.echo.GET("/refresh", app.endpoint.RefreshTokens)
	//app.echo.GET("/secret", app.endpoint.DeleteUser, app.endpoint.ValidateJWT)

	return app, nil
}

func (a *App) Run() error {
	fmt.Println("Server Runnig")

	err := a.echo.Start(":3500")
	if err != nil {
		log.Println(errors.New("Error Start Service"))
	}
	return nil

}
