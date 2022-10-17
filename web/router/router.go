package router

import (
	_ "blog-console/docs"
	"blog-console/web/controller"
	"blog-console/web/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
