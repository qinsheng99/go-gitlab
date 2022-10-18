package controller

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-gitlab/common"
	"github.com/qinsheng99/go-gitlab/gitlab"
)

func (b *Base) GetFile(c *gin.Context) {
	t := c.GetHeader("PRIVATE-TOKEN")
	id := c.Query("id")

	filePath := url.QueryEscape(c.Query("file"))

	file, err := b.cli.GetFile(&gitlab.GetFile{Ref: c.Query("ref"), BaseFile: gitlab.BaseFile{Id: id, Token: t, File: filePath}})
	if err != nil {
		common.Err(c, err)
		return
	}

	c.JSON(http.StatusOK, file)
}

func (b *Base) UploadFile(c *gin.Context) {
	t := c.GetHeader("PRIVATE-TOKEN")
	id := c.Query("id")

	path := ""
	do, err := b.cli.UploadFile(&gitlab.UploadFile{
		BaseFile: gitlab.BaseFile{Token: t, Id: id}, Typ: ".zip", Sha: "", Path: "",
	}, path)
	if err != nil {
		common.Err(c, err)
		return
	}
	//c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=demo.zip")
	c.Header("Content-Transfer-Encoding", "binary")
	//c.File(path)

	c.DataFromReader(http.StatusOK, do.ContentLength, "application/octet-stream", do.Body, nil)

	//defer os.Remove(path)

	//c.JSON(http.StatusOK, "success")
}

func (b *Base) PostFile(c *gin.Context) {
	var p gitlab.PostFile
	var pf *gitlab.RFile
	err := c.ShouldBindJSON(&p)
	if err != nil {
		common.Err(c, err)
		return
	}
	p.Id = c.Query("id")
	p.File = url.QueryEscape(c.Query("file"))
	pf, err = b.cli.PostFile(&p)
	if err != nil {
		common.Err(c, err)
		return
	}
	common.Success(c, pf)
}

func (b *Base) DeleteFile(c *gin.Context) {
	var p gitlab.DeleteFile
	err := c.ShouldBindJSON(&p)
	if err != nil {
		common.Err(c, err)
		return
	}
	p.Id = c.Query("id")
	p.File = url.QueryEscape(c.Query("file"))
	err = b.cli.DeleteFile(&p)
	if err != nil {
		common.Err(c, err)
		return
	}
	common.Success(c, "success")
}

func (b *Base) PutFile(c *gin.Context) {
	var p gitlab.PutFile
	var pf *gitlab.RFile
	err := c.ShouldBindJSON(&p)
	if err != nil {
		common.Err(c, err)
		return
	}
	p.Id = c.Query("id")
	p.File = url.QueryEscape(c.Query("file"))
	pf, err = b.cli.PutFile(&p)
	if err != nil {
		common.Err(c, err)
		return
	}
	common.Success(c, pf)
}

func (b *Base) GetFileRaw(c *gin.Context) {
	t := c.GetHeader("PRIVATE-TOKEN")
	id := c.Query("id")

	filePath := url.QueryEscape(c.Query("file"))

	file, err := b.cli.GetFileRaw(&gitlab.GetFile{Ref: c.Query("ref"), BaseFile: gitlab.BaseFile{Id: id, Token: t, File: filePath}})
	if err != nil {
		common.Err(c, err)
		return
	}

	c.JSON(http.StatusOK, string(file))
}
