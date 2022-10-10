package main

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-gitlab/gitlab"
	"golang.org/x/oauth2"
)

var token func() []byte

func init() {
	token = func() []byte {
		bys, err := os.ReadFile("token")
		if err != nil {
			return nil
		}

		return bys
	}
}

func main() {
	cli := gitlab.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: string(token())})))
	r := gin.Default()

	r.GET("/get-commit", func(c *gin.Context) {
		t := c.GetHeader("PRIVATE-TOKEN")
		id := c.Query("id")
		sha := c.Query("sha")
		var (
			repoCommit interface{}
			err        error
		)
		if len(sha) == 0 {
			repoCommit, err = cli.GetRepoCommit(&gitlab.BaseCommit{Token: t, Id: id})
		} else {
			repoCommit, err = cli.GetOneRepoCommit(&gitlab.CommitOne{BaseCommit: gitlab.BaseCommit{Token: t, Id: id}, Sha: sha})
		}

		if err != nil {
			Err(c, err)
			return
		}

		c.JSON(http.StatusOK, repoCommit)
	})

	r.POST("/post-commit", func(c *gin.Context) {
		t := c.GetHeader("PRIVATE-TOKEN")
		id := c.Query("id")

		repoCommit, err := cli.PostRepoCommit(&gitlab.PostCommit{
			BaseCommit: gitlab.BaseCommit{Token: t, Id: id}, Branch: "demo1", Message: "create demo",
			Actions: []gitlab.Actions{{Action: "create", FilePath: "demo1.txt", Content: "oo", Encoding: "text"}},
		})

		if err != nil {
			Err(c, err)
			return
		}

		c.JSON(http.StatusOK, repoCommit)
	})

	r.GET("/get-tree", func(c *gin.Context) {
		t := c.GetHeader("PRIVATE-TOKEN")
		id := c.Query("id")
		trees, err := cli.GetTree(&gitlab.GetTree{
			GetFile: gitlab.GetFile{Ref: c.Query("ref"), BaseFile: gitlab.BaseFile{Id: id, Token: t}}, Path: c.Query("path"),
		})
		if err != nil {
			Err(c, err)
			return
		}

		c.JSON(http.StatusOK, trees)
	})

	r.GET("/get-file", func(c *gin.Context) {
		t := c.GetHeader("PRIVATE-TOKEN")
		id := c.Query("id")

		filePath := url.QueryEscape(c.Query("file"))

		file, err := cli.GetFile(&gitlab.GetFile{Ref: c.Query("ref"), BaseFile: gitlab.BaseFile{Id: id, Token: t, File: filePath}})
		if err != nil {
			Err(c, err)
			return
		}

		c.JSON(http.StatusOK, file)
	})

	r.GET("/upload-file", func(c *gin.Context) {
		t := c.GetHeader("PRIVATE-TOKEN")
		id := c.Query("id")

		path := "demo.zip"
		err := cli.UploadFile(&gitlab.UploadFile{
			BaseFile: gitlab.BaseFile{Token: t, Id: id}, Typ: ".zip", Sha: "", Path: "",
		}, path)
		if err != nil {
			Err(c, err)
			return
		}
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+path)
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(path)

		//defer os.Remove(path)

		//c.JSON(http.StatusOK, "success")
	})

	r.POST("/post-file", func(c *gin.Context) {
		var p gitlab.PostFile
		var pf *gitlab.RFile
		err := c.ShouldBindJSON(&p)
		if err != nil {
			Err(c, err)
			return
		}
		p.Id = c.Query("id")
		p.File = url.QueryEscape(c.Query("file"))
		pf, err = cli.PostFile(&p)
		if err != nil {
			Err(c, err)
			return
		}
		Success(c, pf)
	})

	r.DELETE("/delete-file", func(c *gin.Context) {
		var p gitlab.DeleteFile
		err := c.ShouldBindJSON(&p)
		if err != nil {
			Err(c, err)
			return
		}
		p.Id = c.Query("id")
		p.File = url.QueryEscape(c.Query("file"))
		err = cli.DeleteFile(&p)
		if err != nil {
			Err(c, err)
			return
		}
		Success(c, "success")
	})

	r.PUT("/put-file", func(c *gin.Context) {
		var p gitlab.PutFile
		var pf *gitlab.RFile
		err := c.ShouldBindJSON(&p)
		if err != nil {
			Err(c, err)
			return
		}
		p.Id = c.Query("id")
		p.File = url.QueryEscape(c.Query("file"))
		pf, err = cli.PutFile(&p)
		if err != nil {
			Err(c, err)
			return
		}
		Success(c, pf)
	})

	r.GET("/get-file-raw", func(c *gin.Context) {
		t := c.GetHeader("PRIVATE-TOKEN")
		id := c.Query("id")

		filePath := url.QueryEscape(c.Query("file"))

		file, err := cli.GetFileRaw(&gitlab.GetFile{Ref: c.Query("ref"), BaseFile: gitlab.BaseFile{Id: id, Token: t, File: filePath}})
		if err != nil {
			Err(c, err)
			return
		}

		c.JSON(http.StatusOK, string(file))
	})

	_ = r.Run(":8080")
}

func Err(c *gin.Context, err error) {
	c.JSON(0, gin.H{
		"error": err.Error(),
	})
	return
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
	return
}
