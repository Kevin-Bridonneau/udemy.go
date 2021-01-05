package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	f, _ := os.Create("golang.html")
	defer f.Close()

	io.Copy(f, resp.Body)
}
