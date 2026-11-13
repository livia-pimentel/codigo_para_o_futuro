package main

import (
	"github.com/colibriproject-dev/colibri-sdk-go"
	"github.com/colibriproject-dev/colibri-sdk-go/pkg/database/sqlDB"
	"github.com/colibriproject-dev/colibri-sdk-go/pkg/web/restserver"
	"github.com/livia-pimentel/codigo-para-o-futuro/src/application/controllers"
)

func init() {
	colibri.InitializeApp()
	sqlDB.Initialize()
	// storage.Initialize()
	// messaging.Initialize()
	// cacheDB.Initialize()
}

func main() {
	restserver.AddRoutes(controllers.NewOpportunityV1Controller().Routes())
	restserver.ListenAndServe()
}