package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetMessages(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(dbh.GetMessages(params["id"], params["mid"]))
}

func init() {
	getMessages := rh.Route{
		"getMessages",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guild/{id}/messages",
		true,
		true,
		runGetMessages,
	}

	rh.Register(getMessages)
}
