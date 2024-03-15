package controller

import (
	"fmt"
	"go-blogs/common/stripmd"
	"go-blogs/models"
	"go-blogs/service"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

var postService service.PostService
var channelService service.ChannelService

// post list
func ListPost(c *gin.Context) {
	plist := postService.GetPostList()
	gintemplate.HTML(c, http.StatusOK, "post/list", gin.H{"plist": plist})
}

// view post
func ViewPost(c *gin.Context) {
	sid, r := c.GetQuery("id")
	var post models.Post
	if r {
		id, _ := strconv.Atoi(sid)
		post = postService.GetPost(id)
	}
	clist := channelService.GetChannelList()
	gintemplate.HTML(c, http.StatusOK, "post/view", gin.H{"post": post, "clist": clist})
}

// delete post
func DeletePost(c *gin.Context) {
	sid, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(sid)
	postService.DelPost(id)
	c.Redirect(http.StatusFound, "/admin/post/list")
}

// post detail 前端
func PostDetail(c *gin.Context) {
	sid := c.Param("id")
	id := cast.ToInt((sid))
	post := postService.GetPost(id)

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	md := []byte(post.Content)
	md = markdown.NormalizeNewlines(md)
	html := markdown.ToHTML(md, parser, nil)

	gintemplate.HTML(c, http.StatusOK, "/post/detail", gin.H{"post": post, "content": template.HTML(html)})
}

// 上传封面
func UploadThumbnails(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "没有文件",
		})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	// The file is received, so let's save it
	pwd, _ := os.Getwd() // 项目根路径
	filePath := fmt.Sprintf("%s/static/thumbnails/%s", pwd, newFileName)
	relativeFilePath := fmt.Sprintf("/static/thumbnails/%s", newFileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "不能保存文件",
		})
		fmt.Printf("err: %v\n", err)
		return
	}
	c.String(http.StatusOK, relativeFilePath)

}

// 添加或者更新
func SavePost(c *gin.Context) {
	var post models.Post

	// id := c.PostForm("id")
	title := c.PostForm("title")
	thumbnail := c.PostForm("thumbnail")
	tags := c.PostForm("tags")
	content := c.PostForm("content")
	channelId := c.PostForm("channelId")

	// 摘要，去掉markdown格式
	summary := stripmd.Strip(content)
	l := len(summary)
	if l >= 200 {
		summary = summary[0:200]
	} else {
		summary = summary[0:l]
	}

	post.Title = title
	post.Thumbnail = thumbnail
	post.Tags = tags
	post.Content = content
	post.Summary = summary
	post.AuthorId = 1 // 实现注册登录后，动态获得
	post.ChannelId = cast.ToInt(channelId)

	if err := c.ShouldBind(&post); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	id := cast.ToInt(c.PostForm("id"))

	if id != 0 {
		postService.UpdatePost(post)
	} else {
		postService.AddPost(post)
	}
	c.Redirect(http.StatusFound, "/admin/post/list")
}
