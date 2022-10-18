package controller

import (
	"github.com/qinsheng99/go-gitlab/gitlab"
)

type Base struct {
	cli gitlab.Inter
}

func NewBase(cli gitlab.Inter) *Base {
	return &Base{cli: cli}
}
