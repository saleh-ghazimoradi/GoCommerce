package utils

type StandardHttpResponse struct {
	error
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

const (
	Ok              = "OK"
	NotValidData    = "Request is not Valid!"
	ProblemInSystem = "Problem in system!"
)
