package main

import (
	"fmt"
	"os"

	"udemy.go/dictionary.go/dictionary"
)

func main() {
	d, err := dictionary.New("./badger")
	handlErr(err)
	defer d.Close()

	d.Add("goland", " a wonderful language")
	entry, _ := d.Get("golang")
	fmt.Println(entry)
}

func handlErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}
