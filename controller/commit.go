package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-gitlab/common"
	"github.com/qinsheng99/go-gitlab/gitlab"
)

func (b *Base) GetCommit(c *gin.Context) {
	t := c.GetHeader("PRIVATE-TOKEN")
	id := c.Query("id")
	sha := c.Query("sha")
	var (
		repoCommit interface{}
		err        error
	)
	if len(sha) == 0 {
		repoCommit, err = b.cli.GetRepoCommit(&gitlab.BaseCommit{Token: t, Id: id})
	} else {
		repoCommit, err = b.cli.GetOneRepoCommit(&gitlab.CommitOne{BaseCommit: gitlab.BaseCommit{Token: t, Id: id}, Sha: sha})
	}

	if err != nil {
		common.Err(c, err)
		return
	}

	c.JSON(http.StatusOK, repoCommit)
}

func (b *Base) PostCommit(c *gin.Context) {
	t := c.GetHeader("PRIVATE-TOKEN")
	id := c.Query("id")

	repoCommit, err := b.cli.PostRepoCommit(&gitlab.PostCommit{
		BaseCommit: gitlab.BaseCommit{Token: t, Id: id}, Branch: "demo1", Message: "create demo",
		Actions: []gitlab.Actions{{Action: "create", FilePath: "demo1.txt", Content: "oo", Encoding: "text"}},
	})

	if err != nil {
		common.Err(c, err)
		return
	}

	c.JSON(http.StatusOK, repoCommit)
}
