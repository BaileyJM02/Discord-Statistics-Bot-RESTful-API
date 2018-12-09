package main

import (
	"fmt"
	"net/http"

	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
	_ "github.com/BaileyJM02/Hue-API/pkg/routes"
	_ "github.com/BaileyJM02/Hue-API/pkg/routes/db"
	"github.com/gorilla/mux"
)

func main() {
	// dbh.People = append(dbh.People, dbh.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &dbh.Address{City: "City X", State: "State X"}})
	// dbh.People = append(dbh.People, dbh.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &dbh.Address{City: "City Z", State: "State Y"}})
	// dbh.People = append(dbh.People, dbh.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
	router := mux.NewRouter()
	// Register routes.
	for route := range rh.Routes {
		router.HandleFunc(rh.Routes[route].Path, rh.Routes[route].Run).Methods(rh.Routes[route].Method)
	}
	fmt.Printf("API is now running with %v/%v endpoints.\n", len(rh.Routes), rh.TotalRoutes)
	http.ListenAndServe(fmt.Sprintf(":%v", "8000"), router)
}
