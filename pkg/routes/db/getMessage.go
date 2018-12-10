package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dbh.GetMessage(params["id"], params["mid"]))
}

func init() {
	getMessage := rh.Route{
		"getMessage",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guild/{id}/message/{mid}",
		true,
		true,
		runGetMessage,
	}

	rh.Register(getMessage)
}
