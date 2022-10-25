package gitlab

import (
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

type Inter interface {
	RepoCommit
	RepoFile
	Graphql
	Project
	GetTree(*GetTree) ([]*Tree, error)
	RetCli() *retryablehttp.Client
}

type RepoCommit interface {
	GetRepoCommit(*BaseCommit) ([]*GetCommit, error)
	PostRepoCommit(*PostCommit) (*GetCommit, error)
	GetOneRepoCommit(*CommitOne) (*GetCommit, error)
}

type RepoFile interface {
	GetFile(*GetFile) (*File, error)
	GetFileRaw(*GetFile) ([]byte, error)
	UploadFile(*UploadFile, string) (*http.Response, error)
	PostFile(*PostFile) (*RFile, error)
	DeleteFile(*DeleteFile) error
	PutFile(*PutFile) (*RFile, error)
}

type Graphql interface {
	Graphql(string, string, string) (*RespRepoGraphql, error)
	GetFileSizeForGraphql(string, string, string) (*RespFileGraphql, error)
}

type Project interface {
	ForkProject(*ForkProjectOption, string) (*ForkProject, error)
	PutProject(*EditProjectOptions, string) (*ForkProject, error)
	CreateProject(*CreateProjectOptions) (*ForkProject, error)
}
