package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func(s*Service)GetFile(c *gin.Context)(int,interface{}){
	code:=c.PostForm("code")
	fmt.Println(code)
	path:= "../data/"+code+"/"

	if _,err:=os.Stat(path);err==nil {
		return s.makeErrJSON(403,40301,"code repeat")
	}else {
		err= os.MkdirAll(path,os.ModePerm)
		if err!=nil {
			return s.makeErrJSON(500,50003,"create dir err")
		}
	}

	file1, err := c.FormFile("file1")
	if err != nil {
		return s.makeErrJSON(500,50001,"file1"+err.Error())
	}
	err = c.SaveUploadedFile(file1, "../data/"+code+"/"+file1.Filename)
	if err!=nil {
		return s.makeErrJSON(500,50002,"file1"+err.Error())
	}

	file2, err := c.FormFile("file2")
	if err != nil {
		return s.makeErrJSON(500,50001,"file2"+err.Error())
	}
	err = c.SaveUploadedFile(file2, "../data/"+code+"/"+file2.Filename)
	if err!=nil {
		return s.makeErrJSON(500,50002,"file2"+err.Error())
	}

	file3, err := c.FormFile("file3")
	if err != nil {
		return s.makeErrJSON(500,50001,"file3"+err.Error())
	}
	err = c.SaveUploadedFile(file3, "../data/"+code+"/"+file3.Filename)
	if err!=nil {
		return s.makeErrJSON(500,50002,"file3"+err.Error())
	}

	file4, err := c.FormFile("file4")
	if err != nil {
		return s.makeErrJSON(500,50001,"file4"+err.Error())
	}
	err = c.SaveUploadedFile(file4, "../data/"+code+"/"+file4.Filename)
	if err!=nil {
		return s.makeErrJSON(500,50002,"file4"+err.Error())
	}

	fmt.Println(file1.Filename,file2.Filename,file3.Filename,file4.Filename)

	return s.makeSuccessJSON(map[string]interface{}{
		"file1":file1,
		"file2":file2,
		"file3":file3,
		"file4":file4,
	})
}
