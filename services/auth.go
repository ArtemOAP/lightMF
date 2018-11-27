package services

import (
	"net/http"
	"log"
	"../config"
)

type MyHandler func(w http.ResponseWriter, r *http.Request)

func FrontController(pass MyHandler) MyHandler {

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



func GetOnly(h MyHandler) MyHandler {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)
			return
		}
		http.Error(w, "get only", http.StatusMethodNotAllowed)
	}
}

func PostOnly(h MyHandler) MyHandler {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			h(w, r)
			return
		}
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}
