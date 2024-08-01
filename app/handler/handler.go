package handler

import (
	"net/http"
	"golang-vercel/app/liveurls"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

// @Tags        Welcome
// @Summary     Hello User
// @Description Endpoint to Welcome user and say Hello "Name"
// @Param       name query string true "Name in the URL param"
// @Accept      json
// @Produce     json
// @Success     200 {object} object "success"
// @Failure     400 {object} object "Request Error or parameter missing"
// @Failure     404 {object} object "When user not found"
// @Failure     500 {object} object "Server Error"
// @Router      /hello/:name [GET]
func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello %v", c.Param("name"))
}
func Itv(c *gin.Context) {
	//path := c.Param("path")
	rid := c.Param("rid")
	ts := c.Query("ts")
	itvobj := &liveurls.Itv{}
	cdn := c.Query("cdn")
	if ts == "" {
	    itvobj.HandleMainRequest(c, cdn, rid)
	} else {
	    itvobj.HandleTsRequest(c, ts)
	}
}





		
