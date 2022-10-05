package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Message string `json:"message"`
	Host    string `json:"host"`
}

func main() {

	app := echo.New()
	app.GET("/", welcome)
	app.GET("/health", health)
	app.Logger.Fatal(app.Start(":8080"))

}

func welcome(c echo.Context) error {
	return c.JSON(http.StatusOK, &response{
		Message: "root path",
		Host:    c.Request().Host,
	})
}

func health(c echo.Context) error {
	log.Println("### HEALTH CHECK HEADER")
	log.Println(c.Request().Header)
	return c.JSON(http.StatusOK, &response{
		Message: "health path",
		Host:    c.Request().Host,
	})
}
