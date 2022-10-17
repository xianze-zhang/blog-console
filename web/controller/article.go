package controller

import (
	"blog-console/common"
	"blog-console/response"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) GetArticle(c *gin.Context) {
	response.NewResponse(c).ToErrorResponse(common.ServerError)
	return
}

// GetArticleList
// @Summary 获取文章列表
// @Tags  article
// @Produce  json
// @Param name query string false "文章名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} common.Error "请求错误"
// @Failure 500 {object} common.Error "内部错误"
// @Router /articles [get]
func (a Article) GetArticleList(c *gin.Context) {

}

// CreateArticle
// @Summary 新增文章
// @Tags  article
// @Produce  json
// @Param name body string true "文章名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} common.Error "请求错误"
// @Failure 500 {object} common.Error "内部错误"
// @Router /articles [post]
func (a Article) CreateArticle(c *gin.Context) {

}

// UpdateArticle
// @Summary 更新文章
// @Tags  article
// @Produce  json
// @Param id path int true "文章 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Article "成功"
// @Failure 400 {object} common.Error "请求错误"
// @Failure 500 {object} common.Error "内部错误"
// @Router /articles/{:id} [put]
func (a Article) UpdateArticle(c *gin.Context) {

}

// UpdateArticleState
// @Summary 更新文章状态
// @Tags  article
// @Produce  json
// @Param id path int true "文章 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Article "成功"
// @Failure 400 {object} common.Error "请求错误"
// @Failure 500 {object} common.Error "内部错误"
// @Router /articles/{:id}/state [put]
func (a Article) UpdateArticleState(c *gin.Context) {

}

// DeleteArticle
// @Summary 删除文章
// @Tags  article
// @Produce  json
// @Param id path int true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} common.Error "请求错误"
// @Failure 500 {object} common.Error "内部错误"
// @Router /articles/{:id} [delete]
func (a Article) DeleteArticle(c *gin.Context) {

}
