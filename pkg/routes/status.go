package routes

import (
	"net/http"

	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
)

func runStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func init() {
	status := rh.Route{
		"e",
		"/db/people",
		"Shhhh",
		"Database",
		"GET",
		"/",
		true,
		true,
		runStatus,
	}

	rh.Register(status)
}
