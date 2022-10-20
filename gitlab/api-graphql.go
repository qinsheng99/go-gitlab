package gitlab

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) Graphql(fullpath, path, ref string) (*RespRepoGraphql, error) {
	var (
		err error
		do  *http.Response
		bys []byte
	)
	body := `
	{
		"query":"query {project(fullPath: \"%s\") {repository {tree(ref: \"%s\", recursive: true, path: \"%s\") {blobs {nodes {name path type}}}}}}"
	}
`
	head := map[string]string{
		//"PRIVATE-TOKEN": f.Token,
		"Content-Type": "application/json",
	}
	do, err = c.request(context.Background(), "POST", "https://gitlab.com/api/graphql", head, strings.NewReader(fmt.Sprintf(body, fullpath, ref, path)), nil)
	if err != nil {
		return nil, err
	}

	bys, err = c.valication(do)
	if err != nil {
		return nil, err
	}
	var resp RespRepoGraphql
	err = json.Unmarshal(bys, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetFileSizeForGraphql(fullpath, path, ref string) (*RespFileGraphql, error) {
	var (
		err error
		do  *http.Response
		bys []byte
	)
	body := `
	{
    "query":"query {project(fullPath: \"%s\"){repository {blobs(ref:\"%s\", paths: [%v]) {nodes{rawSize name path}}}}}"
	}
`
	head := map[string]string{
		//"PRIVATE-TOKEN": f.Token,
		"Content-Type": "application/json",
	}
	do, err = c.request(context.Background(), "POST", "https://gitlab.com/api/graphql", head, strings.NewReader(fmt.Sprintf(body, fullpath, ref, path)), nil)
	if err != nil {
		return nil, err
	}

	bys, err = c.valication(do)
	if err != nil {
		return nil, err
	}
	var resp RespFileGraphql
	err = json.Unmarshal(bys, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
