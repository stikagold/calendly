package Responses

import (
	"calendly/common/models"
	"calendly/repositories/Customer/Elements"
	"encoding/json"
)

type OkResponse struct {
	Answer string
}

func (rsp *OkResponse) ToJson() ([]byte, error) {
	return json.Marshal(rsp)
}

type PacksListResponse struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Data    []models.Pack `json:"data"`
}

func (rsp *PacksListResponse) ToJson() ([]byte, error) {
	return json.Marshal(rsp)
}

type OrderResponse struct {
	Status  bool                    `json:"status"`
	Message string                  `json:"message"`
	Data    []Elements.OrderElement `json:"data"`
}

func (rsp *OrderResponse) ToJson() ([]byte, error) {
	return json.Marshal(rsp)
}
