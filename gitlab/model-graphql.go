package gitlab

type (
	RespRepoGraphql struct {
		Data pr
	}

	pr struct {
		Project repository `json:"project"`
	}

	repository struct {
		Repo tree `json:"repository"`
	}

	tree struct {
		Tree blobs `json:"tree"`
	}

	blobs struct {
		Blobs nodes `json:"blobs"`
	}

	nodes struct {
		Nodes []node `json:"nodes"`
	}

	node struct {
		Name, Path, Type string
	}
)

func (r *RespRepoGraphql) GetNodes() []node {
	return r.Data.Project.Repo.Tree.Blobs.Nodes
}

type (
	RespFileGraphql struct {
		Data struct {
			Project struct {
				Repository struct {
					Blobs struct {
						Nodes struct {
							Node []node
						} `json:"nodes"`
					} `json:"blobs"`
				} `json:"repository"`
			} `json:"project"`
		} `json:"data"`
	}
)
