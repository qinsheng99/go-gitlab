package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-gitlab/gitlab"
	"github.com/qinsheng99/go-gitlab/route"
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

	route.Route(r, cli)

	_ = r.Run(":8080")
}
