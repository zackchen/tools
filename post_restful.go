package main

import (
	"bytes"
	"github.com/deepglint/muses/util/io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		log.Println("Usage: post url action filename")
		return
	}

	url := args[1]
	action := args[2]
	fileName := args[3]

	fileContent, _ := io.ReadBytes(fileName)

	fileReader := bytes.NewReader(fileContent)

	client := &http.Client{}

	request, _ := http.NewRequest("POST", url+"/waitress/"+action, fileReader)

	request.SetBasicAuth("1", "1")
	response, _ := client.Do(request)

	body, _ := ioutil.ReadAll(response.Body)

	log.Println("Response: ", string(body))
}
