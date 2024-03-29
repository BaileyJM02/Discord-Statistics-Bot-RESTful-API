package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	"github.com/BaileyJM02/Hue-API/pkg/logger"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runRemoveGuild(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dbh.RemoveGuild(params["id"])
	guildsdbh, err := dbh.GetGuilds()
	if err != nil {
		logger.Error(err)
	}
	json.NewEncoder(w).Encode(guildsdbh)
}

func init() {
	removeGuild := rh.Route{
		"removeGuild",
		"/db/people",
		"Shhhh",
		"Database",
		"DELETE",
		"/db/guild/{id}",
		true,
		true,
		runRemoveGuild,
	}

	rh.Register(removeGuild)
}
