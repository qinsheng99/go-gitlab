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
	Action          string `json:"action"`
	FilePath        string `json:"file_path"`
	PreviousPath    string `json:"previousPath"`
	Content         string `json:"content"`
	Encoding        string `json:"encoding"`
	LastCommitId    string `json:"last_commit_id"`
	ExecuteFilemode string `json:"execute_filemode"`
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
	Branch  string    `description:"分支" required:"true"`
	Message string    `description:"commit信息" required:"true"`
	Actions []Actions `description:"提交文件" required:"true"`
}
