package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	"github.com/BaileyJM02/Hue-API/pkg/logger"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messagedbh, err := dbh.GetMessage(params["id"], params["mid"])
	if err != nil {
		logger.Error(err)
	}
	json.NewEncoder(w).Encode(messagedbh)
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
