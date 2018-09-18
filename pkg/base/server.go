package base

import (
	"net/http"
)

type CrudServer interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type ListerServer interface {
	List(w http.ResponseWriter, r *http.Request)
}


