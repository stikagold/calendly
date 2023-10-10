package architypes

import (
	"calendly/common/interfaces"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct {
	Core   interfaces.ICore
	Router *mux.Router
}

func (c *Controller) WriteResponse(w http.ResponseWriter, response interfaces.Jsonable, code int) error {
	rpsb, err := response.ToJson()
	if code == 0 {
		code = http.StatusOK
	}
	w.WriteHeader(code)
	_, err = w.Write(rpsb)
	return err
}
