package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

var ApiUrl string

func main() {
	out, err := exec.Command("pngpaste", "-").Output()
	if err != nil {
		log.Println("pngpaste returned non-zero code!")
		log.Println(err)
		return
	}

	res, err := http.Post(ApiUrl, "image/png", bytes.NewReader(out))
	if err != nil {
		log.Println("POST to the API failed!")
		log.Println(err)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("could not read url!")
		log.Println(err)
		return
	}

	fmt.Printf("%s\n", string(body))
}