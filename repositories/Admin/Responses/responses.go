package Responses

import "encoding/json"

type OkResponse struct {
	Answer string
}

func (rsp *OkResponse) ToJson() ([]byte, error) {
	return json.Marshal(rsp)
}
