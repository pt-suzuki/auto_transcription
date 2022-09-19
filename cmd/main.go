package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/pt-suzuki/auto_transcription/infrastructure/server"
)

func main() {
	echo := server.CreateEcho()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	_ = echo.Start(":8080")
}
