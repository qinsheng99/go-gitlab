package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinsheng99/go-gitlab/common"
	"github.com/qinsheng99/go-gitlab/gitlab"
)

func (b *Base) ForkProject(c *gin.Context) {
	var req gitlab.ForkProjectOption
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		common.Err(c, err)
		return
	}

	pro, err := b.cli.ForkProject(&req, c.Query("id"))
	if err != nil {
		common.Err(c, err)
		return
	}
	common.Success(c, pro.ID)
}

func (b *Base) PutProject(c *gin.Context) {
	var req gitlab.EditProjectOptions
	req.Description = common.String(c.Query("des"))

	pro, err := b.cli.PutProject(&req, c.Query("id"))
	if err != nil {
		common.Err(c, err)
		return
	}
	common.Success(c, pro.ID)
}

func (b *Base) CreateProject(c *gin.Context) {
	var req gitlab.CreateProjectOptions
	req.Description = common.String(c.Query("des"))
	req.Name = common.String(c.Query("name"))
	req.Path = common.String(c.Query("name"))
	v := gitlab.PrivateVisibility
	req.Visibility = &v

	pro, err := b.cli.CreateProject(&req)
	if err != nil {
		common.Err(c, err)
		return
	}
	common.Success(c, pro.ID)
}
