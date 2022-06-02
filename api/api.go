package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ramin0x53/addressparser/parse"
)

type request struct {
	Address string `json:"address"`
}

func handleReq(c *gin.Context) {
	var req request
	c.BindJSON(&req)
	c.JSON(200, parse.AddrParse(req.Address))
}

func Run() {
	router := gin.Default()
	router.POST("/", handleReq)
	router.Run()
}
