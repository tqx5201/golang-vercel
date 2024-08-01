package handler

import (
	"net/http"
	"Golang/liveurls"
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

func duanyan(adurl string, realurl any) string {
	var liveurl string
	if str, ok := realurl.(string); ok {
		liveurl = str
	} else {
		liveurl = adurl
	}
	return liveurl
}

func Itv(c *gin.Context) {
	path := c.Param("path")
		rid := c.Param("rid")
		ts := c.Query("ts")
		switch path {
		case "itv":
			if enableTV {
				itvobj := &liveurls.Itv{}
				cdn := c.Query("cdn")
				if ts == "" {
					itvobj.HandleMainRequest(c, cdn, rid)
				} else {
					itvobj.HandleTsRequest(c, ts)
				}
			} else {
				c.String(http.StatusForbidden, "公共服务不提供TV直播")
			}
		case "ysptp":
			if enableTV {
				ysptpobj := &liveurls.Ysptp{}
				if ts == "" {
					ysptpobj.HandleMainRequest(c, rid)
				} else {
					ysptpobj.HandleTsRequest(c, ts, c.Query("wsTime"))
				}
			} else {
				c.String(http.StatusForbidden, "公共服务不提供TV直播")
			}
		case "douyin":
			douyinobj := &liveurls.Douyin{}
			douyinobj.Rid = rid
			douyinobj.Stream = c.DefaultQuery("stream", "flv")
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, douyinobj.GetDouYinUrl()))
		case "douyu":
			douyuobj := &liveurls.Douyu{}
			douyuobj.Rid = rid
			douyuobj.Stream_type = c.DefaultQuery("stream", "flv")
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, douyuobj.GetRealUrl()))
		case "huya":
			huyaobj := &liveurls.Huya{}
			huyaobj.Rid = rid
			huyaobj.Cdn = c.DefaultQuery("cdn", "hwcdn")
			huyaobj.Media = c.DefaultQuery("media", "flv")
			huyaobj.Type = c.DefaultQuery("type", "nodisplay")
			if huyaobj.Type == "display" {
				c.JSON(200, huyaobj.GetLiveUrl())
			} else {
				c.Redirect(http.StatusMovedPermanently, duanyan(adurl, huyaobj.GetLiveUrl()))
			}
		case "bilibili":
			biliobj := &liveurls.BiliBili{}
			biliobj.Rid = rid
			biliobj.Platform = c.DefaultQuery("platform", "web")
			biliobj.Quality = c.DefaultQuery("quality", "10000")
			biliobj.Line = c.DefaultQuery("line", "first")
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, biliobj.GetPlayUrl()))
		case "youtube":
			ytbObj := &liveurls.Youtube{}
			ytbObj.Rid = rid
			ytbObj.Quality = c.DefaultQuery("quality", "1080")
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, ytbObj.GetLiveUrl()))
		case "yy":
			yyObj := &liveurls.Yy{}
			yyObj.Rid = rid
			yyObj.Quality = c.DefaultQuery("quality", "4")
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, yyObj.GetLiveUrl()))
		}
}





		
