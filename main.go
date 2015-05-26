package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"log"
	"os"
)

var base64str string

func init() {
	flag.StringVar(&base64str, "b", "", "base64 encoded text")
}

func main() {
	flag.Parse()
	var out bytes.Buffer
	data1, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		log.Println(err)
		return
	}
	json.Indent(&out, []byte(data1), "", "\t")
	out.WriteTo(os.Stdout)
}
