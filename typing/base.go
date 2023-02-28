package typing

type Home struct {
	Status      bool   `json:"status"`
	Version     string `json:"version"`
	Description string `json:"description"`
}
type Health struct {
	Status string `json:"status`
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type M map[string]interface{}
