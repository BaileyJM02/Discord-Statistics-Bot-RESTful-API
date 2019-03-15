package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	"github.com/BaileyJM02/Hue-API/pkg/logger"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/gorilla/mux"
)

func runGetMessagesTH(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messagesdbh, err := dbh.GetMessagesTH(params["id"])
	if err != nil {
		logger.Error(err)
	}
	json.NewEncoder(w).Encode(messagesdbh)
}

func init() {
	getMessagesTH := rh.Route{
		"getMessagesTH",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/db/guild/{id}/messages/thishour",
		true,
		true,
		runGetMessagesTH,
	}

	rh.Register(getMessagesTH)
}
