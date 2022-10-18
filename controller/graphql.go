package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-gitlab/common"
	"github.com/qinsheng99/go-gitlab/gitlab"
)

func (b *Base) Graphql(c *gin.Context) {
	ref := c.Query("ref")
	data, err := b.cli.Graphql(c.Query("fullpath"), c.Query("path"), ref)
	if err != nil {
		common.Err(c, err)
		return
	}
	var acts []gitlab.Actions
	for _, node := range data.GetNodes() {
		if node.Type == "blob" {
			acts = append(acts, gitlab.Actions{
				Action:   "delete",
				FilePath: node.Path,
			})
		}

	}

	commit, err := b.cli.PostRepoCommit(&gitlab.PostCommit{
		BaseCommit: gitlab.BaseCommit{Token: "", Id: c.Query("id")}, Branch: ref, Message: "delete base",
		Actions: acts,
	})
	if err != nil {
		common.Err(c, err)
		return
	}

	common.Success(c, commit)
}
