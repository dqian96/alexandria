package server

import (
	"encoding/json"
	"log"
	"net/http"

	a "github.com/dqian96/alexandria/archive"
	d "github.com/dqian96/alexandria/director"
	"github.com/gorilla/mux"
)

type valueMsg struct {
	Value string `json:"value"`
}

type errorMsg struct {
	Error  string `json:"errorMessage"`
	Leader string `json:"leader"`
}

func archiveGetHandler(reqID string, w http.ResponseWriter, r *http.Request, d d.Director) {
	key := mux.Vars(r)["key"]
	log.Printf("[%s] Proposal to get value for key '%s' to Director", reqID, key)
	value, in, err := d.Get(reqID, key)
	if err != nil {
		log.Printf("[%s] Proposal unsuccessful: %v", reqID, err)
		//		log.Printf("[%s] Failed to get key '%s': %v", reqID, key, err)
		writeInternalServiceError(w, err, d)
		return
	}
	log.Printf("[%s] Proposal successful: %v", reqID, err)
	if !in {
		log.Printf("[%s] Key '%s' does not exist", reqID, key)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res := valueMsg{Value: value}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	log.Printf("[%s] Get successful!", reqID)
}

func archivePutHandler(reqID string, w http.ResponseWriter, r *http.Request, d d.Director) {
	var valueMsg valueMsg
	key, err := mux.Vars(r)["key"], json.NewDecoder(r.Body).Decode(&valueMsg)
	defer r.Body.Close()
	if err != nil {
		log.Printf("[%s] Could not parse JSON...", reqID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("[%s] Recieved JSON %s", reqID, r.Body)
	log.Printf("[%s] Proposing to put entry (%s, %s) to Director", reqID, key, valueMsg.Value)
	err = d.Put(reqID, a.Entry{
		Key:   key,
		Value: valueMsg.Value,
	})
	if err != nil {
		log.Printf("[%s] Proposal unsuccessful: %v", reqID, err)
		writeInternalServiceError(w, err, d)
		return
	}
	log.Printf("[%s] Proposal successful!", reqID)
	w.WriteHeader(http.StatusCreated)
	log.Printf("[%s] Put successful!", reqID)
}

func archiveDeleteHandler(reqID string, w http.ResponseWriter, r *http.Request, d d.Director) {
	key := mux.Vars(r)["key"]
	log.Printf("[%s] Proposing to delete key '%s' to Director", reqID, key)
	in, err := d.Delete(reqID, key)
	if err != nil {
		log.Printf("[%s] Proposal unsuccessful: %v", reqID, err)
		writeInternalServiceError(w, err, d)
		return
	}
	log.Printf("[%s] Proposal successful!", reqID)
	if !in {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("[%s] Key '%s' does not exist in Archive", reqID, key)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Printf("[%s] Delete successful!", reqID)
}

func writeInternalServiceError(w http.ResponseWriter, err error, d d.Director) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorMsg{
		Error:  err.Error(),
		Leader: d.GetLeader(),
	})
}
