package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	"github.com/BaileyJM02/Hue-API/pkg/logger"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetMembersTD(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	membersdbh, err := dbh.GetMembersTD(params["id"])
	if err != nil {
		logger.Error(err)
	}
	json.NewEncoder(w).Encode(membersdbh)
}

func init() {
	getMembersTD := rh.Route{
		"getMembersTD",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guild/{id}/members/thisday",
		true,
		true,
		runGetMembersTD,
	}

	rh.Register(getMembersTD)
}
