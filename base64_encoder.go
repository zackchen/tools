package main

import (
	"encoding/base64"
	"flag"
	"github.com/deepglint/muses/util/io"
	"log"
	"os"
)

var (
	inputFilePath  string
	outputFilePath string
	method         string
)

func main() {
	flag.StringVar(&method, "method", "d", "Decode(d/D) or Endode(e/E)")
	flag.StringVar(&inputFilePath, "input", "", "Input file path")
	flag.StringVar(&outputFilePath, "output", "", "Output file path")

	flag.Parse()

	if len(os.Args) != 4 {
		log.Println("Args :", len(os.Args))
		flag.Usage()
		return
	}

	log.Println("Method : ", method)

	// if method != "d" || method != "D" || method != "e" || method != "E" {
	// 	flag.Usage()
	// 	return
	// }

	inputContent, err := io.ReadBytes(inputFilePath)

	outputContent := make([]byte, 0)

	if err != nil {
		log.Printf("Read input file %s failed: %s \n", inputFilePath, err)
		return
	}
	if method == "e" || method == "E" {
		base64.StdEncoding.Encode(outputContent, inputContent)
	} else {
		base64.StdEncoding.Decode(outputContent, inputContent)
	}

	log.Println("Output to: ", string(outputContent))

	if err = io.CreateFileByString(outputFilePath, string(outputContent)); err != nil {
		log.Printf("Save output file %s failed: %s \n", outputFilePath, err)
		return
	}

}
