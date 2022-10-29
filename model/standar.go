package model

type ReponseStandar struct {
	Status  bool        `json:"status"`
	Message string      `jsom:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error"`
}
