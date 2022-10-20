package gitlab

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

func (c *Client) ForkProject(options *ForkProjectOption, id string) (*ForkProject, error) {

	var (
		do  *http.Response
		err error
		bys []byte
	)
	baseurl := c.url + "projects/:id/fork"
	baseurl = strings.ReplaceAll(baseurl, ":id", id)

	bys, err = json.Marshal(options)
	if err != nil {
		return nil, err
	}

	do, err = c.request(context.Background(), "POST", baseurl, map[string]string{}, bys, nil)
	if err != nil {
		return nil, err
	}

	bys, err = c.valication(do)
	if err != nil {
		return nil, err
	}

	var project ForkProject
	err = json.Unmarshal(bys, &project)
	if err != nil {

	}
	return &project, nil
}

func (c *Client) PutProject(options *EditProjectOptions, id string) (*ForkProject, error) {
	var (
		do  *http.Response
		err error
		bys []byte
	)
	baseurl := c.url + "projects/:id"
	baseurl = strings.ReplaceAll(baseurl, ":id", id)

	bys, err = json.Marshal(options)
	if err != nil {
		return nil, err
	}

	do, err = c.request(context.Background(), "PUT", baseurl, map[string]string{"Content-Type": "application/json"}, bys, nil)
	if err != nil {
		return nil, err
	}

	bys, err = c.valication(do)
	if err != nil {
		return nil, err
	}

	var project ForkProject
	err = json.Unmarshal(bys, &project)
	if err != nil {

	}
	return &project, nil
}

func (c *Client) CreateProject(options *CreateProjectOptions) (*ForkProject, error) {
	var (
		do  *http.Response
		err error
		bys []byte
	)
	baseurl := c.url + "projects"

	bys, err = json.Marshal(options)
	if err != nil {
		return nil, err
	}

	do, err = c.request(context.Background(), "POST", baseurl, map[string]string{"Content-Type": "application/json"}, bys, nil)
	if err != nil {
		return nil, err
	}

	bys, err = c.valication(do)
	if err != nil {
		return nil, err
	}

	var project ForkProject
	err = json.Unmarshal(bys, &project)
	if err != nil {

	}
	return &project, nil
}
