package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	"github.com/BaileyJM02/Hue-API/pkg/logger"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetMembersTM(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	membersdbh, err := dbh.GetMembersTM(params["id"])
	if err != nil {
		logger.Error(err)
	}
	json.NewEncoder(w).Encode(membersdbh)
}

func init() {
	getMembersTM := rh.Route{
		"getMembersTM",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guild/{id}/members/thisminute",
		true,
		true,
		runGetMembersTM,
	}

	rh.Register(getMembersTM)
}
