package base

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CrudServer defines method to handle CRUD requests
type CrudServer interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

// ListerServer defines method to list items
type ListerServer interface {
	List(w http.ResponseWriter, r *http.Request)
}

type Server struct{}

func (s Server) JsonBody(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-type", "application/json")
	encodedBody, err := json.MarshalIndent(body, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(encodedBody))
}
