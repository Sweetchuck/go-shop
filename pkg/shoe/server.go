package shoe

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/base"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/shoe/model"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/shoe/storage"
	"net/http"
	"strconv"
)

type Server struct {
	base    base.Server
	Storage storage.Handler
}

func (s Server) Create(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)

		return
	}

	var p *model.Shoe
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), 400)

		return
	}

	p.ID = 0

	p2, _ := s.Storage.Insert(*p)
	body := map[string]interface{}{
		"error": "",
		"items": []model.Shoe{
			p2,
		},
	}

	s.base.JsonBody(w, body)
}

func (s Server) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	body := map[string]interface{}{
		"error": "",
		"items": s.Storage.Read(uint(id)),
	}

	s.base.JsonBody(w, body)
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

	body := map[string]interface{}{
		"error": "",
		"new":   "",
	}

	pNew, err := s.Storage.Update(pOld, fields)
	if err != nil {
		w.WriteHeader(403)
		body["err"] = err.Error()
	} else {
		body["new"] = pNew
	}

	s.base.JsonBody(w, body)
}

func (s Server) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	logrus.Info("DELETE id=", id)

	s.Storage.Delete(uint(id))

	body := map[string]interface{}{
		"error": "",
	}

	s.base.JsonBody(w, body)
}

func (s Server) List(w http.ResponseWriter, r *http.Request) {
	body := map[string]interface{}{
		"error": "",
		"items": s.Storage.List(),
	}

	s.base.JsonBody(w, body)
}
