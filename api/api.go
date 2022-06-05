package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramin0x53/addressparser/parse"
	"github.com/thoas/stats"
)

type request struct {
	Address string `json:"address"`
}

// Stats provides response time, status code count, etc.
var Stats = stats.New()

func handleReq(c *gin.Context) {
	var req request

	err := c.BindJSON(&req)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, parse.AddrParse(req.Address))
}

func Run() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			beginning, recorder := Stats.Begin(c.Writer)
			c.Next()
			Stats.End(beginning, stats.WithRecorder(recorder))
		}
	}())

	router.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, Stats.Data())
	})

	router.POST("/", handleReq)

	err := router.Run(":80")
	if err != nil {
		log.Println(err)
	}

}
