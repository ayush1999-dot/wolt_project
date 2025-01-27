package app

import "github.com/gin-gonic/gin"

type App struct {
	Router          *gin.Engine
	ApplicationName string
}

func New(ApplicationName string) *App {
	return &App{
		Router:          gin.Default(),
		ApplicationName: ApplicationName,
	}
}
