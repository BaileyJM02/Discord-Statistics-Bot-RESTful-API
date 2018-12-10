package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetMembers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dbh.GetMembers(params["id"]))
}

func init() {
	getMembers := rh.Route{
		"getMembers",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guild/{id}/members",
		true,
		true,
		runGetMembers,
	}

	rh.Register(getMembers)
}
