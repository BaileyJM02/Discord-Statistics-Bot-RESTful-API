package main

import (
	"fmt"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	_ "github.com/BaileyJM02/Hue-API/pkg/routes"
	_ "github.com/BaileyJM02/Hue-API/pkg/routes/db"

	"github.com/BaileyJM02/Hue-API/pkg/logger"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Register routes.
	for route := range rh.Routes {
		router.HandleFunc(rh.Routes[route].Path, rh.Routes[route].Run).Methods(rh.Routes[route].Method)
	}
	logger.Info(fmt.Sprintf("API is now running with %v/%v endpoints.", len(rh.Routes), rh.TotalRoutes))
	go dbh.Start()
	http.ListenAndServe(fmt.Sprintf(":%v", "8000"), router)
}
