package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range dbh.People {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&dbh.Person{})
}

func init() {
	getPerson := rh.Route{
		"d",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/people/{id}",
		true,
		true,
		runGetPerson,
	}

	rh.Register(getPerson)
}
