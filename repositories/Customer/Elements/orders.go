package Elements

import "calendly/common/models"

type OrderElement struct {
	Count int         `json:"count"`
	Pack  models.Pack `json:"pack"`
}
