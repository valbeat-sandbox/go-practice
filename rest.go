package main

import (
	"net/http"
	"log"
	"github.com/ant0ine/go-json-rest/rest"
)

// http request.
type postIdRequest struct {
	ID int
}
// http response.
type postIdResponse struct {
	Name string
}
// Action method.
func getUser(w *rest.ResponseWriter, req *rest.Request) {
	request := postIdRequest{}
	err := req.DecodeJsonPayload(&request)
	// server error.
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// validation error.
	if request.ID < 1 {
		rest.Error(w, "ID is required", http.StatusBadRequest)
	}
	log.Printf("%#v", request)

	w.WriteJson(&postIdResponse{
		"Gopher",
	})
}


// request handler
func main() {
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		rest.Route{"POST", "/user", getUser},
	)
	log.Printf("Server started.")
	http.ListenAndServe(":10000", &handler)
}
