package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
)

func rungetStats(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dbh.GetStats())
}

func init() {
	getStats := rh.Route{
		"getStats",
		"/stats",
		"Shhhh",
		"Database",
		"GET",
		"/stats",
		true,
		true,
		rungetStats,
	}

	rh.Register(getStats)
}
