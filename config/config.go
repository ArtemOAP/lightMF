package config

import (
	"../entities"
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)
var CONF *entities.Configured
func init(){
	fmt.Println("config init")
	if CONF == nil{
		CONF = &entities.Configured{
			Session: make(map[string]string)}

		source, err := ioutil.ReadFile("./config/config.yml")
		if err != nil {
			log.Fatalf("File config not found!\n")
		}
		errY := yaml.Unmarshal(source, CONF)
		if errY != nil {
			log.Fatalf("File config no valid - error: %v", errY)
		}
		log.Println("start: ",CONF.AppName)




	}
}
func GetConf() *entities.Configured{
	return CONF
}
func SetSesion(key string, val string) {
	CONF.Session[key] = val
}
func GetSesion(key string) string {
	if val, ok := CONF.Session[key]; ok {
		return val
	}
	return ""
}
func HasSesion(key string) bool {
	_, ok := CONF.Session[key]
	return ok
}

func DeleteSession(key string)  {
	delete(CONF.Session, key)
}

