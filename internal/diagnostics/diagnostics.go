package diagnostics

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewDiagnostics() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/healthz", healthz)
	router.HandleFunc("/ready", ready)

	return router
}

func healthz(w http.ResponseWriter, r *http.Request) {
	log.Print("HEALTHZ handler called")
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

func ready(w http.ResponseWriter, r *http.Request) {
	log.Print("READY handler called")
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
