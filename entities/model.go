package entities

import (
	"net/http"
	"time"
)

type User struct {
	Id       int     `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Phone   string  `json:"phone"`
	Email   string  `json:"email"`
	Position string  `json:"position"`
	Salary  int `json:"salary"`
	Path     string  `json:"path,omitempty"`
}
/*
total -- the total number of records available
per_page -- the number of records in each page (page size)
current_page -- the current page number of this data chunk
last_page -- the last page number of this data
next_page_url -- URL of the next page
prev_page_url -- URL of the previous page
from -- the start record of this page, in relation to page size
to -- the end record of this page, in relation to page size
*/

type List struct {
	Total int `json:"total"`
	PerPage int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage int `json:"last_page"`
	NextPageUrl string `json:"next_page_url"`
	PrevPageUrl string `json:"prev_page_url"`
	From int `json:"from"`
	To int `json:"to"`
	Filter string `json:"filter"`
	Users []*User  `json:"users"`
}

type Configured struct {
	AppName string `yaml:"app_name"`
	Session map[string]string
	ListenHostPort string `yaml:"app_listen_host"`
	AppGetWayHost string `yaml:"app_get_way_host"`
	Db struct{
		DbName string `yaml:"db_name"`
		DbUser string `yaml:"db_user"`
		DbPass string `yaml:"db_pass"`
	} `yaml:"db"`
	AdminName string `yaml:"admin_name"`
	AdminPass string `yaml:"admin_pass"`
	TokenLen int `yaml:"token_len"`
	SessionTokenName string `yaml:"session_token_name"`
	SessionTokenMin time.Duration `yaml:"session_token_min"`
}

func (Configured) Init() http.Handler  {
	//TODO init configure
	//fmt.Println("logger init")
	return nil
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

