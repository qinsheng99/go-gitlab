package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/qinsheng99/go-gitlab/controller"
	"github.com/qinsheng99/go-gitlab/gitlab"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	fmt.Println(r.Header)
	return true
}}

func Route(r *gin.Engine, cli gitlab.Inter) {
	base := controller.NewBase(cli)
	func(b *controller.Base) {
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

		r.GET("/socket", func(c *gin.Context) {
			fmt.Println(c.Request.URL.Query())
			// Upgrade our raw HTTP connection to a websocket based one
			conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				log.Print("Error during connection upgradation:", err)
				return
			}
			defer conn.Close()

			// The event loop
			for {
				messageType, message, err := conn.ReadMessage()
				if err != nil {
					log.Println("Error during message reading:", err)
					break
				}
				log.Printf("Received: %s", message)
				err = conn.WriteMessage(messageType, message)
				if err != nil {
					log.Println("Error during message writing:", err)
					break
				}
			}
		})
	}(base)

}
