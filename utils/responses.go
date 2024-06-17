package utils

type Response struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

type Error struct {
	Message string `json:"message"`
}
