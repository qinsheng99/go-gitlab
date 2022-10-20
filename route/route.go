package route

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-gitlab/controller"
	"github.com/qinsheng99/go-gitlab/gitlab"
)

func Route(r *gin.Engine, cli gitlab.Inter) {
	base := controller.NewBase(cli)
	func(b *controller.Base) {
		r.UseRawPath = true
		r.GET("/get-commit", b.GetCommit)
		r.POST("/post-commit", b.PostCommit)

		r.GET("/get-tree", b.GetTree)
		r.GET("/get-file", b.GetFile)
		r.GET("/upload-file", b.UploadFile)
		r.POST("/post-file", b.PostFile)
		r.DELETE("/delete-file", b.DeleteFile)
		r.PUT("/put-file", b.PutFile)
		r.GET("/get-file-raw", b.GetFileRaw)

		r.GET("/ql", b.Graphql)

		r.POST("/fork-project", b.ForkProject)
		r.PUT("/put-project", b.PutProject)
		r.GET("/create-project", b.CreateProject)

	}(base)

}
