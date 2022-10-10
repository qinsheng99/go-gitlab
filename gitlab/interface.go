package gitlab

type Inter interface {
	RepoCommit
	RepoFile
	GetTree(*GetTree) ([]*Tree, error)
}

type RepoCommit interface {
	GetRepoCommit(*BaseCommit) ([]*GetCommit, error)
	PostRepoCommit(*PostCommit) (*GetCommit, error)
	GetOneRepoCommit(*CommitOne) (*GetCommit, error)
}

type RepoFile interface {
	GetFile(*GetFile) (*File, error)
	GetFileRaw(f *GetFile) ([]byte, error)
	UploadFile(*UploadFile, string) error
	PostFile(f *PostFile) (*RFile, error)
	DeleteFile(f *DeleteFile) error
	PutFile(f *PutFile) (*RFile, error)
}
