package gitlab

import (
	"net/http"
)

type Inter interface {
	RepoCommit
	RepoFile
	Graphql
	GetTree(*GetTree) ([]*Tree, error)
	RetCli() *http.Client
}

type RepoCommit interface {
	GetRepoCommit(*BaseCommit) ([]*GetCommit, error)
	PostRepoCommit(*PostCommit) (*GetCommit, error)
	GetOneRepoCommit(*CommitOne) (*GetCommit, error)
}

type RepoFile interface {
	GetFile(*GetFile) (*File, error)
	GetFileRaw(f *GetFile) ([]byte, error)
	UploadFile(*UploadFile, string) (*http.Response, error)
	PostFile(f *PostFile) (*RFile, error)
	DeleteFile(f *DeleteFile) error
	PutFile(f *PutFile) (*RFile, error)
}

type Graphql interface {
	Graphql(fullpath, path, ref string) (*RespRepoGraphql, error)
}
