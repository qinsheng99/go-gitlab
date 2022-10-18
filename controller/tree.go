package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-gitlab/common"
	"github.com/qinsheng99/go-gitlab/gitlab"
)

func (b *Base) GetTree(c *gin.Context) {
	t := c.GetHeader("PRIVATE-TOKEN")
	id := c.Query("id")
	trees, err := b.cli.GetTree(&gitlab.GetTree{
		GetFile: gitlab.GetFile{Ref: c.Query("ref"), BaseFile: gitlab.BaseFile{Id: id, Token: t}}, Path: c.Query("path"), Recursive: true,
	})
	if err != nil {
		common.Err(c, err)
		return
	}

	c.JSON(http.StatusOK, trees)
}
