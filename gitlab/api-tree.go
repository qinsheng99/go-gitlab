package gitlab

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	url1 "net/url"
	"strconv"
	"strings"
)

// GetTree GET projects/:id/repository/tree
func (c *Client) GetTree(f *GetTree) ([]*Tree, error) {
	var (
		err  error
		do   *http.Response
		bys  []byte
		data []*Tree
	)
	url := base + "projects/:id/repository/tree"

	url = strings.ReplaceAll(url, ":id", f.Id)

	head := map[string]string{
		//"PRIVATE-TOKEN": f.Token,
	}

	u := url1.Values{}

	if len(f.Ref) > 0 {
		u.Add("ref", f.Ref)
	}
	if len(f.Path) > 0 {
		u.Add("path", f.Path)
	}
	if f.Recursive {
		u.Add("recursive", "true")
	}

	if f.PerPage > 0 {
		u.Add("per_page", strconv.Itoa(int(f.PerPage)))
	}

	do, err = c.request(context.Background(), "GET", url, head, nil, u)
	if err != nil {
		return nil, err
	}

	bys, err = ioutil.ReadAll(do.Body)

	if err != nil {
		return nil, err
	}

	if do.StatusCode >= 300 {
		return nil, errors.New(do.Status + string(bys))
	}

	err = json.Unmarshal(bys, &data)
	if err != nil {
		return nil, err
	}

	for k := range data {
		tree := data[k]
		if tree.Type == "blob" {
			file, fileerr := c.GetFile(&GetFile{
				BaseFile: BaseFile{Token: f.Token, Id: f.Id, File: url1.QueryEscape(tree.Path)},
				Ref:      f.Ref,
			})
			if fileerr != nil {
				continue
			}
			commit, commiterr := c.GetOneRepoCommit(&CommitOne{
				BaseCommit: BaseCommit{Token: f.Token, Id: f.Id},
				Sha:        file.GetLastCommitId(),
			})
			if commiterr != nil {
				continue
			}

			tree.Size = file.GetFileSize()
			tree.Message = commit.GetMessage()
			tree.CommittedDate = commit.GetCommittedDate()
		}
	}
	return data, nil
}
