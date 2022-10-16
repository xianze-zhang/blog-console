package controller

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) UpdateTagState(c *gin.Context) {}
func (t Tag) GetTagList(c *gin.Context)     {}
func (t Tag) CreateTag(c *gin.Context)      {}
func (t Tag) UpdateTag(c *gin.Context)      {}
func (t Tag) DeleteTag(c *gin.Context)      {}
func (t Tag) GetTag(c *gin.Context)         {}
