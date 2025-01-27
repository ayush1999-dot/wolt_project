package main

import (
	"github.com/wolt/DOPC/src/core/app"
	"github.com/wolt/DOPC/src/core/routers"
	"log"
)

func main() {
	application := app.New("DOPC")

	routers.SetupRouters(application.Router)

	log.Fatal(application.Router.Run(":8080"))
}
