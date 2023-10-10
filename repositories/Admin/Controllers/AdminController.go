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

const path = "admin"

type AdminController struct {
	architypes.Controller

	core   interfaces.ICore
	router *mux.Router
}

func (tc *AdminController) Initial(c interfaces.ICore) error {
	tc.core = c
	tc.router = c.GetRouter()
	tc.RegisterRoutes()
	return nil
}

func (tc *AdminController) RegisterRoutes() {
	tc.router.HandleFunc(helpers.GetFormattedUrl(path, "test"), tc.TestGet).Methods("GET")
	tc.router.HandleFunc(helpers.GetFormattedUrl(path, "add"), tc.Add).Methods("GET")
	tc.router.HandleFunc(helpers.GetFormattedUrl(path, "delete/{id}"), tc.Delete).Methods("GET")
}

func (tc *AdminController) TestGet(w http.ResponseWriter, r *http.Request) {
	tmp := Responses.OkResponse{
		Answer: "okay",
	}
	err := tc.WriteResponse(w, &tmp, http.StatusOK)
	if err != nil {
		helpers.HandleError(err)
	}
}

func (tc *AdminController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	packId := params["id"]
	query := `DELETE FROM 'packs' WHERE id=%s`
	queryPrepared := fmt.Sprintf(query, packId)
	_, err := tc.core.GetDB().Exec(queryPrepared)
	if err != nil {
		fmt.Println(err.Error())
	}
	http.Redirect(w, r, "/admin", 302)
}

func (tc *AdminController) Add(w http.ResponseWriter, r *http.Request) {
	query := `INSERT INTO 'packs' ('name', 'size') VALUES ('%s', %s)`
	queryPrepared := fmt.Sprintf(query, r.URL.Query().Get("name"), r.URL.Query().Get("size"))
	_, err := tc.core.GetDB().Exec(queryPrepared)
	if err != nil {
		fmt.Println(err.Error())
	}
	http.Redirect(w, r, "/admin", 302)
}
