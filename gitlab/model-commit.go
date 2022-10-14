package gitlab

type GetCommit struct {
	Id             string        `json:"id"`
	ShortId        string        `json:"short_id"`
	CreatedAt      string        `json:"created_at"`
	ParentIds      []interface{} `json:"parent_ids"`
	Title          string        `json:"title"`
	Message        string        `json:"message"`
	AuthorName     string        `json:"author_name"`
	AuthorEmail    string        `json:"author_email"`
	AuthoredDate   string        `json:"authored_date"`
	CommitterName  string        `json:"committer_name"`
	CommitterEmail string        `json:"committer_email"`
	CommittedDate  string        `json:"committed_date"`
	Trailers       interface{}   `json:"trailers"`
	WebUrl         string        `json:"web_url"`
}

func (g *GetCommit) GetMessage() string {
	return g.Message
}

func (g *GetCommit) GetCommittedDate() string {
	return g.CommittedDate
}

type Actions struct {
	Action          string `json:"action" description:"操作create、delete、move、update、chmod" required:"true"`
	FilePath        string `json:"file_path" description:"文件路径" required:"true"`
	PreviousPath    string `json:"previousPath" description:"被移动文件的原始完整路径, 适用于move" required:"false"`
	Content         string `json:"content" description:"文件内容，除了 delete、chmod 和 move 外其他操作必填" required:"false"`
	Encoding        string `json:"encoding" description:"text 或者 base64。默认值是 text" required:"false"`
	LastCommitId    string `json:"last_commit_id" description:"最后一个已知的文件提交 ID。仅适用于update、move 和 delete 操作" required:"false"`
	ExecuteFilemode bool   `json:"execute_filemode" description:"适用于 chmod 操作,添加可执行标志" required:"false"`
}

type Tree struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Path          string `json:"path"`
	Mode          string `json:"mode"`
	Size          int64  `json:"size"`
	CommittedDate string `json:"committed_date"`
	Message       string `json:"message"`
}

type BaseCommit struct {
	Token string `json:"-"`
	Id    string `json:"-"`
}

type CommitOne struct {
	BaseCommit
	Sha string `description:"commitId" required:"true"`
}

type PostCommit struct {
	BaseCommit
	Branch      string    `json:"branch" description:"分支" required:"true"`
	Message     string    `json:"commit_message" description:"commit信息" required:"true"`
	StartBranch string    `json:"start_branch" description:"新创建的分支派生自的分支名称" required:"false"`
	AuthorEmail string    `json:"author_email,omitempty" description:"提交者邮箱" required:"false"`
	AuthorName  string    `json:"author_name,omitempty" description:"提交者名字" required:"false"`
	Actions     []Actions `json:"actions" description:"提交文件" required:"true"`
}
