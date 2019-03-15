package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	"github.com/BaileyJM02/Hue-API/pkg/logger"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetGuildStats(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	guilddbh, err := dbh.GetGuildStats(params["id"])
	if err != nil {
		logger.Error(err)
	}
	json.NewEncoder(w).Encode(guilddbh)
}

func init() {
	getGuildStats := rh.Route{
		"getGuildStats",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guild/{id}/stats",
		true,
		true,
		runGetGuildStats,
	}

	rh.Register(getGuildStats)
}
