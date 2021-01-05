package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, err := readFile("test.txt")
	if err != nil {
		fmt.Printf("Error while reading file: %v \n", err)
		return
	}
	fmt.Printf("Your file contain : %v\n", dat)
}

func readFile(filename string) (string, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	if len(dat) == 0 {
		return "", fmt.Errorf("Empty Content (filename : %v )", filename)
	}
	return string(dat), nil
}
