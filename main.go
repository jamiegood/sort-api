package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("so far so good")

	router := mux.NewRouter()
	router.HandleFunc("/sorts/", getSorts).Methods("GET")

	http.ListenAndServe(":8001", router)

}
