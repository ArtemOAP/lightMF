package services

import (
	"net/http"
	"log"
	"../config"
)

type handler func(w http.ResponseWriter, r *http.Request)

func FrontController(pass handler) handler {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("request: ",r.URL)

		for i, v := range r.Cookies() {
			log.Println(i, v)
			if v.Name == "session_token" &&  config.HasSesion(v.Value){
				pass(w, r)
				return
			}
		}
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}
}

