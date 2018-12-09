package routes

// import (
// 	"encoding/json"
// 	"net/http"

// 	dbh "github.com/BaileyJM02/Hue-API/pkg/databaseHandler"
// 	rh "github.com/BaileyJM02/Hue-API/pkg/routeHandler"
// 	"github.com/gorilla/mux"
// )

// func runCreatePerson(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var person dbh.Person
// 	_ = json.NewDecoder(r.Body).Decode(&person)
// 	person.ID = params["id"]
// 	dbh.People = append(dbh.People, person)
// 	json.NewEncoder(w).Encode(dbh.People)
// }

// func init() {
// 	createPerson := rh.Route{
// 		"a",
// 		"/db/people",
// 		"Shhhh",
// 		"Database",
// 		"POST",
// 		"/db/people/{id}",
// 		true,
// 		false,
// 		runCreatePerson,
// 	}

// 	rh.Register(createPerson)
// }
