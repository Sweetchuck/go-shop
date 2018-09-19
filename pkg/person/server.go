package person

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/storage"
	"net/http"
	"strconv"
)

type Server struct {
	Storage storage.Handler
}

func (s Server) Create(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)

		return
	}

	var p *model.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	p.ID = 0

	p2, _ := s.Storage.Insert(*p)
	body := map[string]interface{}{
		"error": "",
		"new":   p2,
	}

	s.jsonBody(w, body)
}

func (s Server) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	body := map[string]interface{}{
		"error": "",
		"items": s.Storage.Read(uint(id)),
	}

	s.jsonBody(w, body)
}

func (s Server) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	pOld := s.Storage.Read(uint(id))
	if pOld.ID == 0 {
		w.WriteHeader(404)

		return
	}

	fields := map[string]interface{}{}
	err := json.NewDecoder(r.Body).Decode(&fields)
	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	delete(fields, "ID")
	delete(fields, "Id")
	delete(fields, "id")

	s.Storage.Update(pOld, fields)
}

func (s Server) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	logrus.Info("DELETE id=", id)

	s.Storage.Delete(uint(id))

	body := map[string]interface{}{
		"error": "",
	}

	s.jsonBody(w, body)
}

func (s Server) List(w http.ResponseWriter, r *http.Request) {
	body := map[string]interface{}{
		"error": "",
		"items": s.Storage.List(),
		"count": s.Storage.Count(),
	}

	s.jsonBody(w, body)
}

func (s Server) jsonBody(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-type", "application/json")
	encodedBody, err := json.MarshalIndent(body, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(encodedBody))
}
