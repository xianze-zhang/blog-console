package router

import (
	"blog-console/web/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := controller.NewArticle()
	tag := controller.NewTag()

	tagCtrl := r.Group("/tags")
	{
		tagCtrl.POST("", tag.CreateTag)
		tagCtrl.DELETE("/:id", tag.DeleteTag)
		tagCtrl.PUT("/:id", tag.UpdateTag)
		tagCtrl.PATCH("/:id/state", tag.UpdateTagState)
		tagCtrl.GET("/:id", tag.GetTag)
		tagCtrl.GET("", tag.GetTagList)
	}

	articleCtrl := r.Group("/articles")
	{
		articleCtrl.POST("", article.CreateArticle)
		articleCtrl.DELETE("/:id", article.DeleteArticle)
		articleCtrl.PUT("/:id", article.UpdateArticle)
		articleCtrl.PATCH("/:id/state", article.UpdateArticleState)
		articleCtrl.GET("/:id", article.GetArticle)
		articleCtrl.GET("", article.GetArticleList)
	}

	return r
}
