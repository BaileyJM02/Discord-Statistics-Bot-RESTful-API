package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

func runCreateGuild(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var guild *discordgo.Guild
	_ = json.NewDecoder(r.Body).Decode(&guild)
	dbh.AddGuild(guild)
	json.NewEncoder(w).Encode(dbh.GetGuild(params["id"]))
}

func init() {
	createGuild := rh.Route{
		"createGuild",
		"/db/people",
		"Shhhh",
		"Database",
		"POST",
		"/db/guild/{id}",
		true,
		true,
		runCreateGuild,
	}

	rh.Register(createGuild)
}
