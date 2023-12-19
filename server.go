package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	val, present := os.LookupEnv("PORT")

	if !present {
		fmt.Println("PORT is not defined. Defaulting to :1323")
		val = ":1323"
	} else {
		val = ":" + val
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(val))
}
