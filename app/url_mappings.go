package app

import "github.com/getmiranda/meli-challenge-api/http"

func mapUrls() {
	pHandler := http.MakePingHandler()
	router.GET("/ping", pHandler.Ping)
}
