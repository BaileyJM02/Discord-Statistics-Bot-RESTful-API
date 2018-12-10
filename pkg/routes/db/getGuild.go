package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetGuild(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dbh.GetGuild(params["id"]))
}

func init() {
	getGuild := rh.Route{
		"getGuild",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guild/{id}",
		true,
		true,
		runGetGuild,
	}

	rh.Register(getGuild)
}
