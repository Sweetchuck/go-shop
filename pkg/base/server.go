package base

import (
	"net/http"
)

// CrudServer defines method to handle CRUD requests
type CrudServer interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

// ListServer defines method to list items
type ListerServer interface {
	List(w http.ResponseWriter, r *http.Request)
}
