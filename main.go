package golib

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func TesteLib(echoInstance *echo.Echo) {
	echoInstance.GET("/lib", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"A": "B"})
	})
}

func NovoMetodo() {
	fmt.Println("NOVO METODO")
}
