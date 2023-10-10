package Controllers

import (
	"calendly/common/architypes"
	"calendly/common/helpers"
	"calendly/common/interfaces"
	"calendly/common/models"
	"calendly/repositories/Customer/Elements"
	"calendly/repositories/Customer/Requests"
	"calendly/repositories/Customer/Responses"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"math"
	"net/http"
)

const path = ""

type CustomerController struct {
	architypes.Controller

	core   interfaces.ICore
	router *mux.Router
}

func (tc *CustomerController) Initial(c interfaces.ICore) error {
	tc.core = c
	tc.router = c.GetRouter()
	tc.RegisterRoutes()
	return nil
}

func (tc *CustomerController) RegisterRoutes() {
	tc.router.HandleFunc("/", tc.Index)
	tc.router.HandleFunc("/packs", tc.GetPacks).Methods(http.MethodGet)
	tc.router.HandleFunc(helpers.GetFormattedUrl(path, "order"), tc.Order).Methods("post")

}

func (tc *CustomerController) TestGet(w http.ResponseWriter, r *http.Request) {
	tmp := Responses.OkResponse{
		Answer: "okay",
	}
	err := tc.WriteResponse(w, &tmp, http.StatusOK)
	if err != nil {
		helpers.HandleError(err)
	}
}

func (tc *CustomerController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./public/index.html")
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (tc *CustomerController) GetPacks(w http.ResponseWriter, r *http.Request) {
	var packs []models.Pack
	db := tc.core.GetDB()
	rows, err := db.Query("SELECT * FROM packs")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		pack := models.Pack{}
		err = rows.Scan(&pack.Id, &pack.Name, &pack.Size)
		if err != nil {
			return
		}
		packs = append(packs, pack)
	}

	response := Responses.PacksListResponse{
		Status: true,
		Data:   packs,
	}

	tc.WriteResponse(w, &response, http.StatusOK)
}

func (tc *CustomerController) Order(w http.ResponseWriter, r *http.Request) {
	var req Requests.OrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err.Error())
	}

	var packs []models.Pack
	db := tc.core.GetDB()
	rows, err := db.Query("SELECT * FROM packs ORDER BY size DESC ")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		pack := models.Pack{}
		err = rows.Scan(&pack.Id, &pack.Name, &pack.Size)
		if err != nil {
			return
		}
		packs = append(packs, pack)
	}

	// Start calculating
	requestedCount := req.Count
	var response []Elements.OrderElement
	for _, elem := range packs {
		tmpCount := float64(requestedCount) / float64(elem.Size)
		if tmpCount >= 1 {
			realCount := int(math.Round(tmpCount - 0.5))
			currentOrder := Elements.OrderElement{Count: realCount, Pack: elem}
			response = append(response, currentOrder)
			requestedCount -= realCount * elem.Size
			fmt.Println(elem.Name+"it's possible, going to calculate", realCount)
		}
	}
	if requestedCount > 0 {
		tmp := Elements.OrderElement{
			Count: 1,
			Pack:  packs[len(packs)-1],
		}
		response = append(response, tmp)
	}
	finalResponse := Responses.OrderResponse{
		Status: true,
		Data:   response,
	}
	if len(response) == 0 {
		finalResponse.Status = false
		finalResponse.Message = "Can not confirm order, can not configure packs"
	}
	tc.WriteResponse(w, &finalResponse, http.StatusOK)
}
