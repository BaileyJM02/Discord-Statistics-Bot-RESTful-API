package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

func runCreateMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var message *discordgo.MessageCreate
	_ = json.NewDecoder(r.Body).Decode(&message)
	dbh.AddMessage(message)
	json.NewEncoder(w).Encode(dbh.GetMessage(params["id"], params["mid"]))
}

func init() {
	createMessage := rh.Route{
		"createMessage",
		"/db/people",
		"Shhhh",
		"Database",
		"POST",
		"/db/guild/{id}/message/{mid}",
		true,
		true,
		runCreateMessage,
	}

	rh.Register(createMessage)
}
