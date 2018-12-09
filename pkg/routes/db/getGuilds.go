package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
)

func runGetGuilds(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dbh.GetGuilds())
}

func init() {
	getGuilds := rh.Route{
		"c",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guilds",
		true,
		true,
		runGetGuilds,
	}

	rh.Register(getGuilds)
}
