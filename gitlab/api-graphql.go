package gitlab

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	bys, err = ioutil.ReadAll(do.Body)
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
