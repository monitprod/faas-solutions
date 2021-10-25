package function

import "encoding/json"

type Response struct {
	Message string `json:"message:"`
}

func (r *Response) ToMap() (res map[string]interface{}) {
	a, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(a, &res)
	return
}
