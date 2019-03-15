package routes

import (
	"encoding/json"
	"net/http"

	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
)

func runBackup(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]byte(`{"error":"Depriciated!"}`))
}

func init() {
	backup := rh.Route{
		"backup",
		"/backup",
		"Shhhh",
		"Database",
		"GET",
		"/backup",
		true,
		true,
		runBackup,
	}

	rh.Register(backup)
}
