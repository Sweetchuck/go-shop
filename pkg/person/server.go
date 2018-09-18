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

func (s Server) Insert(w http.ResponseWriter, r *http.Request) {
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
		"items": []model.Person{
			p2,
		},
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
