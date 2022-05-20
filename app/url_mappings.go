package app

import (
	mutants_db "github.com/getmiranda/meli-challenge-api/datasources/postgres/mutants"
	"github.com/getmiranda/meli-challenge-api/http"
	"github.com/getmiranda/meli-challenge-api/repository/db"
	"github.com/getmiranda/meli-challenge-api/services"
)

func mapUrls() {
	pHandler := http.MakePingHandler()
	router.GET("/ping", pHandler.Ping)

	dbClient := mutants_db.GetClient()
	dbRepo := db.MakeDBRepository(dbClient)
	hService := services.MakeHumansService(dbRepo)
	hHandler := http.MakeHumanHandler(hService)
	router.POST("/mutant/", hHandler.IsMutant)
}
