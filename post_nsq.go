package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/deepglint/muses/util/io"
	"log"
	"net/http"
)

var (
	nsqdAddress = flag.String("nsqd_address", "", "Url")
	topic       = flag.String("topic", "", "Topic")
	action      = flag.String("action", "", "Action")
	fileName    = flag.String("file_path", "", "File path")
)

type Command struct {
	SessionId string
	Cmd       string
	Params    map[string]interface{}
}

func main() {

	flag.Parse()

	fileContent, _ := io.ReadBytes(*fileName)

	cmd := &Command{
		SessionId: "abc",
		Cmd:       "waitress/" + *action,
		Params:    make(map[string]interface{}),
	}
	cmd.Params["data"] = []byte(fileContent)

	// log.Println("DATA: ", string(cmd.Params["data"]))

	cmdContent, _ := json.Marshal(cmd)
	log.Println("Cmd content: ", string(cmdContent))
	cmdReader := bytes.NewReader(cmdContent)

	// fileReader := bytes.NewReader(fileContent)

	client := &http.Client{}
	url := "http://" + *nsqdAddress + "/pub?topic=" + *topic
	log.Println("URL: ", url)
	request, _ := http.NewRequest("POST", url, cmdReader)

	response, _ := client.Do(request)

	log.Println("Response: ", response)

}
