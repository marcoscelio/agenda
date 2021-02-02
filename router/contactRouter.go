package router

import (
	"encoding/json"
	"net/http"

	"../dao"
	"../models"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"gopkg.in/mgo.v2/bson"
)

// var contactDao = dao.ContactDao{}

type ContactService struct {
	Dao dao.IContactDao
}

var (
	getReqProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "get_request_total",
		Help: "The total number of 'get' request processed events",
	})

	postReqProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "post_request_total",
		Help: "The total number of 'post' request processed events",
	})

	deleteReqProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "delete_request_total",
		Help: "The total number of 'delete' request processed events",
	})

	putReqProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "put_request_total",
		Help: "The total number of 'put' request processed events",
	})
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (contactService *ContactService) GetAll(w http.ResponseWriter, r *http.Request) {
	defer getReqProcessed.Inc()
	contacts, err := contactService.Dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, contacts)
}

func (contactService *ContactService) GetByID(w http.ResponseWriter, r *http.Request) {
	defer getReqProcessed.Inc()
	params := mux.Vars(r)
	contact, err := contactService.Dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
		return
	}
	respondWithJson(w, http.StatusOK, contact)
}

func (contactService *ContactService) FindByName(w http.ResponseWriter, r *http.Request) {
	defer getReqProcessed.Inc()
	params := mux.Vars(r)
	contacts, err := contactService.Dao.FindByName(params["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid contact Name")
		return
	}
	respondWithJson(w, http.StatusOK, contacts)
}

func (contactService *ContactService) Create(w http.ResponseWriter, r *http.Request) {
	defer postReqProcessed.Inc()
	defer r.Body.Close()
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	contact.ID = bson.NewObjectId()
	if err := contactService.Dao.Create(contact); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, contact)
}

func (contactService *ContactService) Update(w http.ResponseWriter, r *http.Request) {
	defer putReqProcessed.Inc()
	defer r.Body.Close()
	params := mux.Vars(r)
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := contactService.Dao.Update(params["id"], contact); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": contact.Name + " atualizado com sucesso!"})
}

func (contactService *ContactService) Delete(w http.ResponseWriter, r *http.Request) {
	defer deleteReqProcessed.Inc()
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := contactService.Dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
