package gitlab

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	url1 "net/url"
	"os"
	"strings"
)

// GetFile GET projects/:id/repository/files/:file_path
func (c *Client) GetFile(f *GetFile) (*File, error) {
	var (
		do   *http.Response
		err  error
		bys  []byte
		data File
	)
	if err = f.validation(); err != nil {
		return nil, err
	}
	url := base + "projects/:id/repository/files/:file_path"

	url = strings.ReplaceAll(url, ":id", f.Id)
	url = strings.ReplaceAll(url, ":file_path", f.File)

	head := map[string]string{
		//"PRIVATE-TOKEN": f.Token,
	}

	u := url1.Values{}
	if len(f.Ref) > 0 {
		u.Add("ref", f.Ref)
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

	return &data, nil
}

// UploadFile GET projects/:id/repository/archive[.format]
func (c *Client) UploadFile(f *UploadFile, upload string) error {
	var (
		do   *http.Response
		err  error
		file *os.File
	)
	//if err = f.validation(); err != nil {
	//	return err
	//}
	url := base + "projects/:id/repository/archive[.format]"

	url = strings.ReplaceAll(url, ":id", f.Id)
	url = strings.ReplaceAll(url, "[.format]", f.Typ)

	u := url1.Values{}
	if len(f.Sha) != 0 {
		u.Add("sha", f.Sha)
	}
	if len(f.Path) != 0 {
		u.Add("path", f.Path)
	}

	head := map[string]string{
		//"PRIVATE-TOKEN": f.Token,
	}

	do, err = c.request(context.Background(), "GET", url, head, nil, u)
	if err != nil {
		return err
	}

	if do.StatusCode >= 300 {
		return errors.New("")
	}

	file, err = os.Create(upload)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, do.Body)
	if err != nil {
		return err
	}

	return nil
}

// PostFile POST projects/:id/repository/files/:file_path
func (c *Client) PostFile(f *PostFile) (_ *RFile, _ error) {
	var (
		do   *http.Response
		err  error
		bys  []byte
		data RFile
	)
	if err = f.validation(); err != nil {
		return nil, err
	}
	url := base + "projects/:id/repository/files/:file_path"

	url = strings.ReplaceAll(url, ":id", f.Id)
	url = strings.ReplaceAll(url, ":file_path", f.File)

	head := map[string]string{
		//"PRIVATE-TOKEN": f.Token,
		"Content-Type": "application/json",
	}

	bys, err = json.Marshal(f)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bys))

	do, err = c.request(context.Background(), "POST", url, head, bytes.NewReader(bys), nil)

	bys, err = ioutil.ReadAll(do.Body)

	if err != nil {
		return nil, err
	}

	if do.StatusCode >= 300 {
		return nil, errors.New(do.Status + string(bys))
	}

	_ = json.Unmarshal(bys, &data)

	return &data, nil
}

// DeleteFile DELETE projects/:id/repository/files/:file_path
func (c *Client) DeleteFile(f *DeleteFile) error {
	var (
		do  *http.Response
		err error
		bys []byte
	)
	if err = f.validation(); err != nil {
		return err
	}
	url := base + "projects/:id/repository/files/:file_path"

	url = strings.ReplaceAll(url, ":id", f.Id)
	url = strings.ReplaceAll(url, ":file_path", f.File)

	head := map[string]string{
		//"PRIVATE-TOKEN": f.Token,
		"Content-Type": "application/json",
	}

	bys, err = json.Marshal(f)
	if err != nil {
		return err
	}

	do, err = c.request(context.Background(), "DELETE", url, head, bytes.NewReader(bys), nil)

	bys, err = ioutil.ReadAll(do.Body)

	if err != nil {
		return err
	}

	if do.StatusCode >= 300 {
		return errors.New(do.Status + string(bys))
	}

	fmt.Println(string(bys))

	return nil
}

// PutFile PUT projects/:id/repository/files/:file_path
func (c *Client) PutFile(f *PutFile) (*RFile, error) {
	var (
		do   *http.Response
		err  error
		bys  []byte
		data RFile
	)
	if err = f.validation(); err != nil {
		return nil, err
	}
	url := base + "projects/:id/repository/files/:file_path"

	url = strings.ReplaceAll(url, ":id", f.Id)
	url = strings.ReplaceAll(url, ":file_path", f.File)
	head := map[string]string{
		//"PRIVATE-TOKEN": f.Token,
		"Content-Type": "application/json",
	}

	bys, err = json.Marshal(f)
	if err != nil {
		return nil, err
	}

	do, err = c.request(context.Background(), "PUT", url, head, bytes.NewReader(bys), nil)

	bys, err = ioutil.ReadAll(do.Body)

	if err != nil {
		return nil, err
	}

	if do.StatusCode >= 300 {
		return nil, errors.New(do.Status + string(bys))
	}

	_ = json.Unmarshal(bys, &data)

	return &data, nil
}

// GetFileRaw GET projects/:id/repository/files/:file_path/raw
func (c *Client) GetFileRaw(f *GetFile) ([]byte, error) {
	var (
		do  *http.Response
		err error
		bys []byte
	)
	if err = f.validation(); err != nil {
		return nil, err
	}
	url := base + "projects/:id/repository/files/:file_path/raw"

	url = strings.ReplaceAll(url, ":id", f.Id)
	url = strings.ReplaceAll(url, ":file_path", f.File)

	head := map[string]string{
		//"PRIVATE-TOKEN": f.Token,
	}

	u := url1.Values{}
	if len(f.Ref) > 0 {
		u.Add("ref", f.Ref)
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

	return bys, nil
}
