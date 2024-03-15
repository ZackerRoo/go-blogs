package controller

import (
	"fmt"
	"go-blogs/models"
	"go-blogs/service"
	"net/http"
	"strconv"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

var channel service.ChannelService

// channel list
func ListChannel(c *gin.Context) {
	c2 := channel.GetChannelList()
	gintemplate.HTML(c, http.StatusOK, "channel/list", gin.H{"clist": c2})
}

// view channel
func ViewChannel(c *gin.Context) {
	sid, r := c.GetQuery("id")
	var chann models.Channel
	if r {
		id, _ := strconv.Atoi(sid)
		chann = channel.GetChannel(id)
	}
	gintemplate.HTML(c, http.StatusOK, "channel/view", gin.H{"channel": chann})
}

// delete channel
func DeleteChannel(c *gin.Context) {
	sid, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(sid)
	channel.DelChannel(id)
	c.Redirect(http.StatusFound, "/admin/channel/list")
}

// 添加或者更新
func SaveChannel(c *gin.Context) {
	var chann models.Channel
	if err := c.ShouldBind(&chann); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	chann.Status, _ = strconv.Atoi(c.PostForm("status"))

	id, _ := c.GetPostForm("id")

	if id != "0" {
		channel.UpdateChannel(chann)
	} else {
		channel.AddChannel(chann)
	}
	c.Redirect(http.StatusFound, "/admin/channel/list")
}
