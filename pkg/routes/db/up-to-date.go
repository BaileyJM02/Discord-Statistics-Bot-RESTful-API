package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/bwmarrin/discordgo"
)

func runUpToDate(w http.ResponseWriter, r *http.Request) {
	var ready *discordgo.Ready
	_ = json.NewDecoder(r.Body).Decode(&ready)
	dbh.CheckUpToDate(ready)
	json.NewEncoder(w).Encode(dbh.GetGuilds())
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
