package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

type Service struct {
	conf 	conf
	Router 	*gin.Engine
}

type conf struct {
	Server struct{
		Port 	string
		Key 	string
	}
}

func(s *Service) init() {
	s.initConfig()
	s.initRouter()
}

func(s *Service) initConfig(){
	c:= new(conf)
	_,err:= toml.DecodeFile("./config/config.toml",c)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	s.conf= *c
	fmt.Println(*c)
}

func (s *Service) makeErrJSON(httpStatusCode int, errCode int, msg interface{}) (int, interface{}) {
	return httpStatusCode, &gin.H{"error": errCode, "msg": fmt.Sprint(msg)}
}

func (s *Service) makeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, &gin.H{"error": 0, "msg": "success", "data": data}
}