package controller

import (
	"net/http"
	"log"
	"../dataManager"
	"encoding/json"
	"strconv"
	"../entities"
	"../services"
	"../config"
	"time"
	"strings"
	"fmt"
	)

const limit = 50

func List(w http.ResponseWriter, r *http.Request) {

	var(
		er error
		page int
		items []*entities.User
	)
	conf:= config.GetConf()

	if r.URL.Path != "/list/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	page = 0
	 if param:=r.URL.Query().Get("page"); param != ""{
		 page ,er = strconv.Atoi(param)
		 if er != nil{
		 	log.Println(er)
			 page = 0
		 }
	 }
	 filter:= r.URL.Query().Get("filter")
	 count:= dataManager.GetInstance().GetCount(filter)
	 offset:=  page * limit
	 if page == 1 || page == 0{
	 	page= 1
	 	offset = 0
	 }
	//log.Printf("start_select")
	items,er = dataManager.GetInstance().GetRowsWithFiles(limit,offset,filter)
	//log.Printf("end_select")
	if er!= nil{
		http.Error(w,"error db",http.StatusInternalServerError)
	}
	list := new(entities.List)
	list.Users = items
	list.Total = count
	list.PerPage = limit
	list.CurrentPage = page
	list.LastPage =count / limit - 1
	list.Filter = filter
	if page < list.LastPage{
		list.NextPageUrl = conf.AppGetWayHost + "/list/?page=" + strconv.Itoa(page+1)
	}
	if page > 1{
		list.PrevPageUrl = conf.AppGetWayHost + "/list/?page=" + strconv.Itoa(page-1)
	}
	list.To = (list.CurrentPage) * list.PerPage
	list.From = list.PerPage*(list.CurrentPage -1)  + 1
	data, er := json.Marshal(list)
	if er!= nil{
		http.Error(w,"error 1",http.StatusInternalServerError)
	}
	w.Header().Set("Server", "GO")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}
func User(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	log.Println(r.URL)
	p := strings.Split(r.URL.Path, "/")
	fmt.Printf("%v",p)
	id,_:= strconv.Atoi(p[2])


	user,er:=dataManager.GetInstance().GetRowById(id)
	if er!= nil{
		log.Println(er)
	}

	json.NewEncoder(w).Encode(user)


	//http.ServeFile(w, r, "./public/list.html")

}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")


	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./public/list.html")

}


func Index(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")


	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.Redirect(w,r,"/home/",http.StatusSeeOther)

}

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method == "POST" {
		r.ParseForm()
		name:= r.PostForm.Get("name")
		pass:= r.PostForm.Get("password")
		conf:= config.GetConf()
		if name == conf.AdminName && pass == conf.AdminPass{
			sessionToken :=services.GenerateToken(conf.TokenLen)
			config.SetSesion(sessionToken,name)
			http.SetCookie(w, &http.Cookie{
				Name:    conf.SessionTokenName,
				Value:   sessionToken,
				Path: "/",
				Expires: time.Now().Add(conf.SessionTokenMin * time.Minute),
			})


			http.Redirect(w,r,"/home/",http.StatusSeeOther)
			return
		}
		http.Redirect(w,r, conf.AppGetWayHost + "/login/",http.StatusSeeOther)
		return
	}
	http.ServeFile(w, r, "./public/login.html")
}


func Logout(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	for _, v := range r.Cookies() {
		if v.Name == "session_token" &&  config.HasSesion(v.Value){
			config.DeleteSession(v.Value)
			return
		}
	}
	conf:= config.GetConf()
	http.Redirect(w,r, conf.AppGetWayHost + "/login/",http.StatusSeeOther)

}