package function

type EventPayload struct {
	Execution         int64 `json:"execution"`
	UsersPerExecution int64 `json:"usersPerExecution"`
}

type Response struct {
	Message string `json:"message:"`
}
