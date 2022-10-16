package controller

import (
	"blog-console/common"
	"blog-console/respone"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) GetArticle(c *gin.Context) {
	respone.NewResponse(c).ToErrorResponse(common.ServerError)
	return
}
func (a Article) GetArticleList(c *gin.Context) {

}
func (a Article) CreateArticle(c *gin.Context) {

}
func (a Article) UpdateArticle(c *gin.Context) {

}
func (a Article) UpdateArticleState(c *gin.Context) {

}
func (a Article) DeleteArticle(c *gin.Context) {

}
