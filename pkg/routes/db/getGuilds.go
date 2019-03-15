package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	"github.com/BaileyJM02/Hue-API/pkg/logger"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
)

func runGetGuilds(w http.ResponseWriter, r *http.Request) {
	guilddbh, err := dbh.GetGuilds()
	if err != nil {
		logger.Error(err)
	}
	json.NewEncoder(w).Encode(guilddbh)
}

func init() {
	getGuilds := rh.Route{
		"getGuilds",
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
