package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	"github.com/BaileyJM02/Hue-API/pkg/logger"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/bwmarrin/discordgo"
)

func runUpToDate(w http.ResponseWriter, r *http.Request) {
	var ready discordgo.Ready
	_ = json.NewDecoder(r.Body).Decode(&ready)
	dbh.CheckUpToDateReady(ready)
	guilddbh, err := dbh.GetGuilds()
	if err != nil {
		logger.Error(err)
	}
	json.NewEncoder(w).Encode(guilddbh)
}

func init() {
	upToDate := rh.Route{
		"upToDate",
		"/db/people",
		"Shhhh",
		"Database",
		"POST",
		"/db/up-to-date/",
		true,
		true,
		runUpToDate,
	}

	rh.Register(upToDate)
}
