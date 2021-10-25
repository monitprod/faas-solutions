package function

import "encoding/json"

type EventPayload struct {
	Execution         int64 `json:"execution"`
	UsersPerExecution int64 `json:"usersPerExecution"`
}

func (e *EventPayload) ToMap() (res map[string]interface{}) {
	a, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(a, &res)
	return
}

func EventPayloadFromMap(m map[string]interface{}) (*EventPayload, error) {
	e := EventPayload{}

	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, &e)
	}
	return &e, err
}
