package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"./controller"
	"./config"
	"./services"
	)

func main() {

	conf:= config.GetConf()
	fs := http.FileServer(http.Dir("/mnt/volume_lon1_01/zru_photos/img/"))
	http.Handle("/img/", http.StripPrefix("/img/", fs))
	http.Handle("/stat/", http.StripPrefix("/stat", http.FileServer(http.Dir("./public/stat"))))
	http.HandleFunc("/list/",   services.FrontController(controller.List))
	http.HandleFunc("/home/",   services.FrontController(controller.Home))
	http.HandleFunc("/login/",  controller.Login)
	http.HandleFunc("/logout/",  controller.Logout)
	err := http.ListenAndServe(conf.ListenHostPort,nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}





