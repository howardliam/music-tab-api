package utils

type Response struct {
	StatusCode int
	Data       interface{}
	Error      Error
}

type Error struct {
	Message string
}
