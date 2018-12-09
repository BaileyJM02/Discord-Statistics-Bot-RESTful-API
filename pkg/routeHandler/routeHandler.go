package routeHandler

import (
	"net/http"
)

type Route struct {
	Name        string
	Usage       string
	Description string
	Category    string
	Method      string
	Path        string
	Secure      bool
	Enabled     bool
	Run         func(http.ResponseWriter, *http.Request)
}

var (
	Routes      map[string]Route
	TotalRoutes = 0
)

func Register(rut Route) {
	if Routes == nil {
		Routes = make(map[string]Route)
	}

	if rut.Enabled {
		Routes[rut.Name] = rut
	}

	TotalRoutes++
}
