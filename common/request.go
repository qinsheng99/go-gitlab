package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-gitlab/gitlab"
)

type Data interface {
	[]*gitlab.GetCommit | *gitlab.GetCommit | []*gitlab.Tree | *gitlab.File | *gitlab.RFile | ~string
}

func Err(c *gin.Context, err error) {
	c.JSON(0, gin.H{
		"error": err.Error(),
	})
	return
}

func Success[T Data](c *gin.Context, data T) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
	return
}
