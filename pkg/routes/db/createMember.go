package routes

import (
	"encoding/json"
	"net/http"

	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

func runCreateMember(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var member *discordgo.Member
	_ = json.NewDecoder(r.Body).Decode(&member)
	dbh.AddMember(member, member.GuildID)
	memberdbh, _ := dbh.GetMember(params["id"], params["mid"])
	json.NewEncoder(w).Encode(memberdbh)
}

func init() {
	createMember := rh.Route{
		"createMember",
		"/db/people",
		"Shhhh",
		"Database",
		"POST",
		"/db/guild/{id}/member/{mid}",
		true,
		true,
		runCreateMember,
	}

	rh.Register(createMember)
}
