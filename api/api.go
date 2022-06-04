package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ramin0x53/addressparser/parse"
)

type request struct {
	Address string `json:"address"`
}

func handleReq(c *gin.Context) {
	var req request

	err := c.BindJSON(&req)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, parse.AddrParse(req.Address))
}

func Run() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/", handleReq)

	err := router.Run(":80")
	if err != nil {
		log.Println(err)
	}

}
