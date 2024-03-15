package initialize

import (
	"fmt"
	"net/http"

	"go-blogs/common/global"
	"go-blogs/controller"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func Router() {

	engine := gin.Default()

	// 静态资源请求映射
	fmt.Printf("hellllo")
	engine.Static("/assets", "./assets")
	engine.StaticFS("/static", http.Dir("./static"))

	engine.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "templates/frontend",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	})

	// 前台
	engine.GET("/", controller.Index)
	engine.GET("/post/:id", controller.PostDetail)

	mw := gintemplate.NewMiddleware(gintemplate.TemplateConfig{
		Root:         "templates/backend",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	})

	// 后台管理员前端接口
	web := engine.Group("/admin", mw)

	{
		// index
		web.GET("/", controller.AdminIndex)

	}

	{
		// channel
		web.GET("/channel/list", controller.ListChannel)
		web.GET("/channel/view", controller.ViewChannel)
		web.POST("/channel/save", controller.SaveChannel)
		web.GET("/channel/del", controller.DeleteChannel)
	}

	{
		// post
		web.GET("/post/list", controller.ListPost)
		web.GET("/post/view", controller.ViewPost)
		web.POST("/post/save", controller.SavePost)
		web.GET("/post/del", controller.DeletePost)

		web.POST("/post/upload", controller.UploadThumbnails)

	}

	// 启动、监听端口
	s := global.Config.Server
	post := fmt.Sprintf(":%s", s.Port)
	fmt.Printf("post: %v\n", post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}
