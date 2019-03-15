package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runSetPrefix(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var prefix string
	_ = json.NewDecoder(r.Body).Decode(&prefix)
	dbh.SetPrefix(params["id"], prefix)

	prefix2, _ := dbh.GetPrefix(params["id"])

	json.NewEncoder(w).Encode(prefix2)
}

func init() {
	setPrefix := rh.Route{
		"setPrefix",
		"/db/people",
		"Shhhh",
		"Database",
		"POST",
		"/db/guild/{id}/prefix/{prefix}",
		true,
		true,
		runSetPrefix,
	}

	rh.Register(setPrefix)
}
