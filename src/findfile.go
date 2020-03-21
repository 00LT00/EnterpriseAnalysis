package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)
type file struct {
	Time 	time.Time
	Name	string
}

func(s *Service)findFile(c *gin.Context)(int,interface{}){
	files:=new([]file)
	code:= c.Param("code")
	path:="../data/"+code+"/"
	fileInfos,err:=ioutil.ReadDir(path)
	if err != nil {
		return s.makeErrJSON(500,50000,err.Error())
	}
	for _,v:=range fileInfos{
		*files = append(*files,file{
			Time: v.ModTime(),
			Name: v.Name(),
		})
	}

	return s.makeSuccessJSON(*files)
}
