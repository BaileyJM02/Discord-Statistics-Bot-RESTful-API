package routes

// import (
// 	"encoding/json"
// 	"net/http"

// 	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
// 	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
// 	"github.com/gorilla/mux"
// )

// func runDeletePerson(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for index, item := range dbh.People {
// 		if item.ID == params["id"] {
// 			dbh.People = append(dbh.People[:index], dbh.People[index+1:]...)
// 			break
// 		}
// 		json.NewEncoder(w).Encode(dbh.People)
// 	}
// }

// func init() {
// 	deletePerson := rh.Route{
// 		"b",
// 		"/db/people",
// 		"Shhhh",
// 		"Database",
// 		"DELETE",
// 		"/db/people/{id}",
// 		true,
// 		false,
// 		runDeletePerson,
// 	}

// 	rh.Register(deletePerson)
// }
