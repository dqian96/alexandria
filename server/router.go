package server

import (
	"fmt"
	"log"
	"net/http"

	d "github.com/dqian96/alexandria/director"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const (
	archivePath     = "/archive"
	keyResourcePath = archivePath + "/{key}"
)

func setHandlers(router *mux.Router, d d.Director) {
	router.HandleFunc(keyResourcePath, func(w http.ResponseWriter, r *http.Request) {
		uuid, _ := uuid.NewUUID()
		reqID := uuid.String()
		log.Println(fmt.Sprintf("[%s] Recieved request - GET value from key from host '%s'", reqID, r.Host))
		archiveGetHandler(reqID, w, r, d)
	}).Methods("GET")

	router.HandleFunc(keyResourcePath, func(w http.ResponseWriter, r *http.Request) {
		uuid, _ := uuid.NewUUID()
		reqID := uuid.String()
		log.Println(fmt.Sprintf("[%s] Recieved request - PUT value to key from host '%s'", reqID, r.Host))
		archivePutHandler(reqID, w, r, d)
	}).Methods("PUT")

	router.HandleFunc(keyResourcePath, func(w http.ResponseWriter, r *http.Request) {
		uuid, _ := uuid.NewUUID()
		reqID := uuid.String()
		log.Println(fmt.Sprintf("[%s] Recieved request - DELETE key from host '%s'", reqID, r.Host))
		archiveDeleteHandler(reqID, w, r, d)
	}).Methods("DELETE")
}

// NewRouter creates and configures a mux.Router
func NewRouter(d d.Director) *mux.Router {
	router := mux.NewRouter()
	setHandlers(router, d)
	return router
}
