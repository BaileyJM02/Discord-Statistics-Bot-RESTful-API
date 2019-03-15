package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
)

func runCheck(w http.ResponseWriter, r *http.Request) {
	var helper *dbh.CheckHelper
	_ = json.NewDecoder(r.Body).Decode(&helper)
	go dbh.CheckUpToDate(helper.Guild, helper.Members)
	json.NewEncoder(w).Encode("Sent.")

}

func init() {
	check := rh.Route{
		"check",
		"/db/check",
		"Shhhh",
		"Database",
		"POST",
		"/db/check/",
		true,
		true,
		runCheck,
	}

	rh.Register(check)
}
