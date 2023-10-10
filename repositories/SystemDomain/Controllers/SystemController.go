package Controllers

import (
	"calendly/common/architypes"
	"calendly/common/helpers"
	"calendly/common/interfaces"
	"calendly/repositories/Customer/Responses"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const path = "system"

type SystemController struct {
	architypes.Controller

	core   interfaces.ICore
	router *mux.Router
}

func (tc *SystemController) Initial(c interfaces.ICore) error {
	tc.core = c
	tc.router = c.GetRouter()
	tc.RegisterRoutes()
	return nil
}

func (tc *SystemController) RegisterRoutes() {
	tc.router.HandleFunc(fmt.Sprintf("/%s/health", path), tc.TestGet).Methods("GET")
}

func (tc *SystemController) TestGet(w http.ResponseWriter, r *http.Request) {
	tmp := Responses.OkResponse{
		Answer: "okay",
	}
	err := tc.WriteResponse(w, &tmp, http.StatusOK)
	if err != nil {
		helpers.HandleError(err)
	}
}
