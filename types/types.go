package types

import "encoding/json"

type LambdaResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func (l LambdaResponse) String() string {
	data, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(data)
}
