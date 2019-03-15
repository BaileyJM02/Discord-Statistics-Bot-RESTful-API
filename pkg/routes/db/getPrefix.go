package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	"github.com/BaileyJM02/Hue-API/pkg/logger"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetPrefix(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	prefix, err := dbh.GetPrefix(params["id"])
	if err != nil {
		logger.Error(err)
	}
	json.NewEncoder(w).Encode(prefix)
}

func init() {
	getPrefix := rh.Route{
		"getPrefix",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guild/{id}/prefix",
		true,
		true,
		runGetPrefix,
	}

	rh.Register(getPrefix)
}
