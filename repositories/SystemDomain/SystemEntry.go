package SystemDomain

import (
	"calendly/common/helpers"
	"calendly/common/interfaces"
	"calendly/repositories/SystemDomain/Controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Entry struct {
	core   interfaces.ICore
	router *mux.Router

	// Controllers
	tc Controllers.SystemController
}

func (e *Entry) Build(core interfaces.ICore) error {
	fmt.Println("In test domain build function")
	e.core = core
	e.router = core.GetRouter()
	e.router.HandleFunc("/health", e.HealthCheck).Methods("GET")

	err := e.tc.Initial(core)
	helpers.HandleError(err)

	return nil
}

func (e *Entry) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
