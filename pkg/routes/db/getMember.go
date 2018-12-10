package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetMember(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dbh.GetMember(params["id"], params["mid"]))
}

func init() {
	getMember := rh.Route{
		"getMember",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guild/{id}/member/{mid}",
		true,
		true,
		runGetMember,
	}

	rh.Register(getMember)
}
