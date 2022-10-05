package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	jwtcontroller "tchen/jwt-auth/api/module/jwt/controller"
)

var (
	PORT = os.Getenv("PORT")
)

type RouterMux struct {
	router *mux.Router
}

func (r *RouterMux) Start() {
	r.router.HandleFunc("/v1.0/access-token/b2b", jwtcontroller.GetJwtToken).Methods("POST")

	log.Println("Listening on port " + PORT)
	http.ListenAndServe(":"+PORT, r.router)
}
