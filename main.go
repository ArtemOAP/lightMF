package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"./controller"
	"./config"
	"./services"
	"./routs"
		"regexp"
)

func main() {

	conf:= config.GetConf()
	fs := http.FileServer(http.Dir("/mnt/volume_lon1_01/zru_photos/img/"))
	//http.Handle("/img/", http.StripPrefix("/img/", fs))
	//http.Handle("/stat/", http.StripPrefix("/stat", http.FileServer(http.Dir("./public/stat"))))
	//http.HandleFunc("/list/",   services.FrontController(services.GetOnly(controller.List)))
	//http.HandleFunc("/home/",   services.FrontController(services.GetOnly(controller.Home)))
	//http.HandleFunc("/login/",  controller.Login)
	//http.HandleFunc("/logout/",  controller.Logout)

	rh:= new (routs.RegexpHandler)
	rh.DefaultHandle("/img/", http.StripPrefix("/img/", fs))
	rh.DefaultHandle("/stat/", http.StripPrefix("/stat", http.FileServer(http.Dir("./public/stat"))))
	rh.DefaultHandleFunc("/list/",   services.FrontController(services.GetOnly(controller.List)))
	rh.DefaultHandleFunc("/home/",   services.FrontController(services.GetOnly(controller.Home)))
	rh.DefaultHandleFunc("/login/",  controller.Login)
	rh.DefaultHandleFunc("/logout/",  controller.Logout)
	rh.HandleFunc(regexp.MustCompile("/user/[0-9]+"),services.FrontController(services.GetOnly(controller.User)))

	err := http.ListenAndServe(conf.ListenHostPort,rh)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}



