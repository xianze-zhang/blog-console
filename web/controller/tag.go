package controller

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// UpdateTagState
// @Summary 更新标签状态
// @Tags  tag
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} common.Error "请求错误"
// @Failure 500 {object} common.Error "内部错误"
// @Router /tags/{:id}/state [put]
func (t Tag) UpdateTagState(c *gin.Context) {

}

// GetTagList
// @Summary 获取多个标签
// @Tags  tag
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} common.Error "请求错误"
// @Failure 500 {object} common.Error "内部错误"
// @Router /tags [get]
func (t Tag) GetTagList(c *gin.Context) {

}

// CreateTag
// @Summary 新增标签
// @Tags  tag
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} common.Error "请求错误"
// @Failure 500 {object} common.Error "内部错误"
// @Router /tags [post]
func (t Tag) CreateTag(c *gin.Context) {

}

// UpdateTag
// @Summary 更新标签
// @Tags  tag
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} common.Error "请求错误"
// @Failure 500 {object} common.Error "内部错误"
// @Router /tags/{:id} [put]
func (t Tag) UpdateTag(c *gin.Context) {

}

// DeleteTag
// @Summary 删除标签
// @Tags  tag
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} common.Error "请求错误"
// @Failure 500 {object} common.Error "内部错误"
// @Router /tags/{:id} [delete]
func (t Tag) DeleteTag(c *gin.Context) {

}

// GetTag
// @Summary 获取标签详情
// @Tags  tag
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} common.Error "请求错误"
// @Failure 500 {object} common.Error "内部错误"
// @Router /tags/{:id} [get]
func (t Tag) GetTag(c *gin.Context) {

}
