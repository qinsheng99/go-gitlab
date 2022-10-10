package gitlab

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetRepoCommit GET projects/:id/repository/commits
func (c *Client) GetRepoCommit(comm *BaseCommit) ([]*GetCommit, error) {
	var (
		err  error
		do   *http.Response
		bys  []byte
		data []*GetCommit
	)
	url := base + "projects/:id/repository/commits"

	url = strings.ReplaceAll(url, ":id", comm.Id)

	head := map[string]string{
		//"PRIVATE-TOKEN": comm.Token,
	}

	do, err = c.request(context.Background(), "GET", url, head, nil, nil)
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

	return data, nil
}

// GetOneRepoCommit GET projects/:id/repository/commits/:sha
func (c *Client) GetOneRepoCommit(comm *CommitOne) (*GetCommit, error) {
	var (
		err  error
		do   *http.Response
		bys  []byte
		data GetCommit
	)
	url := base + "projects/:id/repository/commits/:sha"

	url = strings.ReplaceAll(url, ":id", comm.Id)
	url = strings.ReplaceAll(url, ":sha", comm.Sha)

	head := map[string]string{
		//"PRIVATE-TOKEN": comm.Token,
	}

	do, err = c.request(context.Background(), "GET", url, head, nil, nil)
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

	return &data, nil
}

// PostRepoCommit Post projects/:id/repository/commits
func (c *Client) PostRepoCommit(f *PostCommit) (*GetCommit, error) {
	var (
		data *GetCommit
		bys  []byte
		err  error
		do   *http.Response
	)
	url := base + "projects/:id/repository/commits"

	url = strings.ReplaceAll(url, ":id", f.Id)

	head := map[string]string{
		//"PRIVATE-TOKEN": f.Token,
		"Content-Type": "application/json",
	}

	body := map[string]interface{}{
		"branch":         f.Branch,
		"commit_message": f.Message,
		"actions":        f.Actions,
	}

	bys, err = json.MarshalIndent(body, "", "   ")
	if err != nil {
		return nil, err
	}

	do, err = c.request(context.Background(), "POST", url, head, bytes.NewReader(bys), nil)
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

	return data, nil
}
