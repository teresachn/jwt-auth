package controller

import (
	"net/http"
)

func GetJwtToken(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
}
