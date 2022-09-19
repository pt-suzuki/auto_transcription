package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pt-suzuki/auto_transcription/infrastructure/firestorage"
	"github.com/pt-suzuki/auto_transcription/infrastructure/firestore"
	echo2 "github.com/pt-suzuki/auto_transcription/src/provider/echo"
	echo3 "github.com/pt-suzuki/auto_transcription/src/router/echo"
)

func CreateEcho() *echo.Echo {
	fireStoreClient := firestore.GetClient()
	fireStorageClient := firestorage.GetClient()

	serviceProvider := echo2.Wire(fireStoreClient, fireStorageClient)

	e := echo3.GetRouter(serviceProvider)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	return e
}
