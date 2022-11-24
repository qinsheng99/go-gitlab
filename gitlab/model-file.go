package gitlab

import (
	"errors"
)

type File struct {
	FileName        string `json:"file_name"`
	FilePath        string `json:"file_path"`
	Size            int64  `json:"size"`
	Encoding        string `json:"encoding"`
	ContentSha256   string `json:"content_sha256"`
	Ref             string `json:"ref"`
	BlobId          string `json:"blob_id"`
	CommitId        string `json:"commit_id"`
	LastCommitId    string `json:"last_commit_id"`
	ExecuteFilemode bool   `json:"execute_filemode"`
	Content         string `json:"content"`
}

func (f *File) GetFileSize() int64 {
	return f.Size
}

func (f *File) GetCommitId() string {
	return f.CommitId
}

func (f *File) GetLastCommitId() string {
	return f.LastCommitId
}

type BaseFile struct {
	Token string `json:"-" description:"token" required:"false"`
	Id    string `json:"-" description:"项目id" required:"true"`
	File  string `json:"-" description:"提交文件路径" required:"true"`
}

type GetFile struct {
	BaseFile
	Ref string `description:"branch, tag or commit" required:"false"`
}

func (g *GetFile) validation() error {
	if len(g.Id) == 0 || len(g.File) == 0 {
		return errors.New("id and file is required")
	}
	return nil
}

type GetTree struct {
	GetFile
	Path      string `description:"目录路径" required:"false"`
	Recursive bool   `description:"递归获取" required:"false"`
	PerPage   int64  `description:"数量" required:"false"`
	Page      int64  `description:"页数" required:"false"`
}

type UploadFile struct {
	BaseFile
	Typ  string `description:"下载文件类型:tar.gz, tar.bz2, tbz, tbz2, tb2, bz2, tar, and zip" required:"true"`
	Sha  string `description:"commitId" required:"false"`
	Path string `description:"下载文件子路径,默认整个库" required:"false"`
}

func (u *UploadFile) validation() error {
	if len(u.Id) == 0 || len(u.File) == 0 || len(u.Typ) == 0 {
		return errors.New("id, file and typ  is required")
	}
	return nil
}

type Files struct {
	Branch        string `json:"branch" description:"提交的分支" required:"true"`
	StartBranch   string `json:"start_branch,omitempty" description:"提交分支的基分支" required:"false"`
	AuthorEmail   string `json:"author_email,omitempty" description:"提交者邮箱" required:"false"`
	AuthorName    string `json:"author_name,omitempty" description:"提交者名字" required:"false"`
	CommitMessage string `json:"commit_message" description:"提交commit信息" required:"true"`
}

type PostFile struct {
	BaseFile
	Files
	Encoding        string `json:"encoding,omitempty" description:"内容格式,默认text,可选base64" required:"false"`
	Content         string `json:"content" description:"文件内容" required:"true"`
	ExecuteFilemode bool   `json:"execute_filemode,omitempty" description:"文件执行flag" required:"false"`
}

func (p *PostFile) validation() error {
	if len(p.Id) == 0 || len(p.File) == 0 || len(p.Branch) == 0 || len(p.Content) == 0 || len(p.CommitMessage) == 0 {
		return errors.New("id, file, content, commit_message and branch  is required")
	}
	return nil
}

type DeleteFile struct {
	BaseFile
	Files
	LastCommitId string `json:"last_commit_id"`
}

func (p *DeleteFile) validation() error {
	if len(p.Id) == 0 || len(p.File) == 0 || len(p.Branch) == 0 || len(p.CommitMessage) == 0 {
		return errors.New("id, file, commit_message and branch  is required")
	}
	return nil
}

type RFile struct {
	File   string `json:"file_path"`
	Branch string `json:"branch"`
}

type PutFile struct {
	PostFile
	LastCommitId string `json:"last_commit_id"`
}
