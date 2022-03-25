# go-go-file

## Description

This project is a wrapper over the gofile service API - gofile.io
The service allows you to upload, store and share files.

## Install

```
go get github.com/mikeyuniverse/go-go-file
```

## Opportunities

1. **Upload file** <br>
Upload a file using the UploadFile method. The method takes the path to the file and returns a link to download the file upon successful upload.
1. **Get account information** <br>
The GetAccountDetails method allows you to check account information. The method prints information to the console.

## Authorization
You need an API key to use it. It can be obtained when authenticating in the service:<br>
1. Go to page - https://gofile.io/myProfile
1. Click - "Login with your email"
1. Receive an email with a login link
1. Follow this link
1. Chapter "My Profile"
1. Copy "Api token"

## Usage example

```go
package main

import (
	"fmt"
	"github.com/mikeyuniverse/go-go-file"
	"log"
)

const TOKEN = "yourToken"

func main() {
	client := gofile.NewClient(TOKEN)

	downloadLink, err := client.UploadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("File uploaded\nDOWNLOAD URL -", downloadLink)

	account, err := client.GetAccountDetails()
	if err != nil {
		log.Fatal(err)
		return
	}
	account.Info()
}
```
