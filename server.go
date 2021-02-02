package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	dbconfig "./config"
	"./dao"
	"./router"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var details = dbconfig.Config{}
var contactDao = &dao.ContactDao{}
var contactRouter = router.ContactService{}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func init() {
	details.Read()
	contactDao.Server = details.Server
	contactDao.Database = details.Database
	contactDao.Connect()
	contactRouter.Dao = contactDao
}

func main() {
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/api/v1/contacts", contactRouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/contacts/{id}", contactRouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/contacts/name/{name}", contactRouter.FindByName).Methods("GET")
	r.HandleFunc("/api/v1/contacts", contactRouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/contacts/{id}", contactRouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/contacts/{id}", contactRouter.Delete).Methods("DELETE")

	var _, portErr = strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if portErr != nil {
		log.Fatal("Missing port number environment variable value")
	}
	if _, err := strconv.Atoi(os.Getenv("PORT")); err != nil {
		log.Fatal("Bad format port number environment variable value")
	}
	var readTimeout, readTimeoutErr = strconv.ParseInt(os.Getenv("READ_TIMEOUT"), 10, 64)
	if readTimeoutErr != nil {
		log.Fatal("Missing read timeout environment variable value READ_TIMEOUT")
	}

	var writeTimeout, writeTimeoutErr = strconv.ParseInt(os.Getenv("WRITE_TIMEOUT"), 10, 64)
	if writeTimeoutErr != nil {
		log.Fatal("Missing write timeout environment variable value WRITE_TIMEOUT")
	}

	// muxWithMiddlewares := http.TimeoutHandler(r, time.Second*time.Duration(timeout), "Timeout!")
	fmt.Println("Server running in port:", os.Getenv("PORT"))

	srv := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      r,
		ReadTimeout:  time.Second * time.Duration(readTimeout),
		WriteTimeout: time.Second * time.Duration(writeTimeout),
	}

	log.Fatal(srv.ListenAndServe())

	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), muxWithMiddlewares))
}
