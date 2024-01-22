package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func startServerWithSubmitter() *mux.Router {
	router := mux.NewRouter()

	submitter := &submitter{}

	router.HandleFunc("/submit", submitter.Submit)

	return router
}

func (s *submitter) Submit(w http.ResponseWriter, r *http.Request) {
	buf := make([]byte, 1000)

	_, err := r.Body.Read(buf)
	if err != nil {
		panic(err)
	}

	tx := string(buf)
	s.queuedTxs = append(s.queuedTxs, tx)

	w.WriteHeader(http.StatusOK)
}
