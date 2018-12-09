package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
)

func runGetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dbh.People)
}

func init() {
	getPeople := rh.Route{
		"c",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/people",
		true,
		true,
		runGetPeople,
	}

	rh.Register(getPeople)
}
