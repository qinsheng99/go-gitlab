package gitlab

type RespRepoGraphql struct {
	Data pr
}

type pr struct {
	Project repository `json:"project"`
}

type repository struct {
	Repo tree `json:"repository"`
}

type tree struct {
	Tree blobs `json:"tree"`
}

type blobs struct {
	Blobs nodes `json:"blobs"`
}

type nodes struct {
	Nodes []node `json:"nodes"`
}

type node struct {
	Name, Path, Type string
}

func (r *RespRepoGraphql) GetNodes() []node {
	return r.Data.Project.Repo.Tree.Blobs.Nodes
}
