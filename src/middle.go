package main

import "github.com/gin-gonic/gin"

func(s *Service)check(c *gin.Context){
	Key:=c.GetHeader("sign")
	if Key != s.conf.Server.Key	{
		c.JSON(s.makeErrJSON(403,40300,"forbidden"))
		c.Abort()
		return
	}
	return
}
