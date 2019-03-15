package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	"github.com/BaileyJM02/Hue-API/pkg/logger"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runRemoveMember(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dbh.RemoveMember(params["id"], params["mid"])
	membersdbh, err := dbh.GetMembers(params["id"])
	if err != nil {
		logger.Error(err)
	}
	json.NewEncoder(w).Encode(membersdbh)
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
