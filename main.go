package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Sort ...
type Sort struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var sorts []Sort

func getSortsHandler(w http.ResponseWriter, r *http.Request) {

	sort := Sort{ID: 1, Title: "hoho", Body: "Yup body here"}

	sorts = append(sorts, sort)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sorts)
}

// func getSort(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	// for _, item := range sorts {
// 	// 	// if item.ID == params["id"] {
// 	// 	// 	json.NewEncoder(w).Encode(item)
// 	// 	// 	break
// 	// 	// }
// 	// 	// return
// 	// }
// 	json.NewEncoder(w).Encode(&Sort{})
// }
func main() {
	fmt.Println("so far so good")

	router := mux.NewRouter()
	router.HandleFunc("/sorts/", getSortsHandler).Methods("GET")

	http.ListenAndServe(":8001", router)

}
