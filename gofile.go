package gofile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	urlGetServer         = "https://api.gofile.io/getServer"
	urlUploadFile        = "https://%s.gofile.io/uploadFile"
	urlGetAccountDetails = "https://api.gofile.io/getAccountDetails?token=%s&allDetails=true"
)

type AccountDetails struct {
	Token              string `json:"token"`
	Email              string `json:"email"`
	Tier               string `json:"tier"`
	RootFolder         string `json:"rootFolder"`
	FoldersCount       int    `json:"foldersCount"`
	FilesCount         int    `json:"filesCount"`
	TotalSize          int    `json:"totalSize"`
	TotalDownloadCount int    `json:"totalDownloadCount"`
}

type Account struct {
	Status string         `json:"status"`
	Data   AccountDetails `json:"data"`
}

func (a *Account) Info() {
	fmt.Printf("\nACCOUNT INFO\nEMAIL - %s\nTIER - %s\nTOTAL DOWNLOADS - %d\nTOTAL FILES - %d\nTOTAL SIZE - %d\n",
		a.Data.Email, a.Data.Tier, a.Data.TotalDownloadCount, a.Data.FilesCount, a.Data.TotalSize)
}

type ResponseUploadFileData struct {
	DownloadPage string `json:"downloadPage"`
	Code         string `json:"code"`
	ParentFolder string `json:"parentFolder"`
	FileId       string `json:"fileId"`
	FileName     string `json:"fileName"`
	Md5          string `json:"md5"`
	DirectLink   string `json:"directLink"`
	Info         string `json:"info"`
}

type ResponseUploadFile struct {
	Status string                 `json:"status"`
	Data   ResponseUploadFileData `json:"data"`
}

type ServerDesc struct {
	Server string `json:"server"`
}

type Server struct {
	Status string     `json:"status"`
	Data   ServerDesc `json:"data"`
}

type Client struct {
	client *http.Client
	token  string
}

func NewClient(apiToken string) *Client {
	return &Client{
		client: &http.Client{},
		token:  apiToken,
	}
}

func (c *Client) getServer() (Server, error) {
	response, err := c.client.Get(urlGetServer)
	if err != nil {
		return Server{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Server{}, err
	}

	var r Server
	if err = json.Unmarshal(body, &r); err != nil {
		return Server{}, err
	}
	return r, nil
}

func (c *Client) UploadFile(filePath string) (string, error) {
	server, err := c.getServer()
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf(urlUploadFile, server.Data.Server)

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	writer.WriteField("token", c.token)

	err = writer.Close()
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, writer.FormDataContentType(), body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var r ResponseUploadFile
	if err = json.Unmarshal(bodyBytes, &r); err != nil {
		return "", err
	}

	return r.Data.DownloadPage, nil
}

func (c *Client) GetAccountDetails() (Account, error) {
	url := fmt.Sprintf(urlGetAccountDetails, c.token)
	response, err := c.client.Get(url)
	if err != nil {
		return Account{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Account{}, err
	}

	var r Account
	if err = json.Unmarshal(body, &r); err != nil {
		return Account{}, err
	}
	return r, nil
}
