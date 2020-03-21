package main

import(
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func(s *Service) initRouter(){
	router := gin.Default()
	router.Use(cors.Default(),s.check)

	router.POST("/upload",
		func(c *gin.Context) {
			c.JSON(s.GetFile(c))
		})
	router.GET("/find/:code",
		func(c *gin.Context) {
			c.JSON(s.findFile(c))
		})
	s.Router = router
	err:=router.Run(s.conf.Server.Port)
	panic(err)
}
