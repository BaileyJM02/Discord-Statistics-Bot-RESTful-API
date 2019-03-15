package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	"github.com/BaileyJM02/Hue-API/pkg/logger"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetMembersTH(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	membersdbh, err := dbh.GetMembersTH(params["id"])
	if err != nil {
		logger.Error(err)
	}
	json.NewEncoder(w).Encode(membersdbh)
}

func init() {
	getMembersTH := rh.Route{
		"getMembersTH",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guild/{id}/members/thishour",
		true,
		true,
		runGetMembersTH,
	}

	rh.Register(getMembersTH)
}
