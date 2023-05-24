package model

type Colecta struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type ColectaResponse struct {
	Message string `json:"message"`
}

type ColectaErrorResponse struct {
	Message string `json:"message"`
}
