package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runRemoveMember(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dbh.RemoveMember(params["id"], params["mid"])
	json.NewEncoder(w).Encode(dbh.GetMembers(params["id"]))
}

func init() {
	removeMember := rh.Route{
		"removeMember",
		"/db/people",
		"Shhhh",
		"Database",
		"DELETE",
		"/db/guild/{id}/member/{mid}",
		true,
		true,
		runRemoveMember,
	}

	rh.Register(removeMember)
}
