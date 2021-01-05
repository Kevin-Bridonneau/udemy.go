package main

import (
	"flag"
	"fmt"
	"os"

	"udemy.go/gencert/cert"
	"udemy.go/gencert/html"
	"udemy.go/gencert/pdf"
)

func main() {
	outputType := flag.String("type", "pdf", "Output type of the certificate")
	csvFile := flag.String("file", "students.csv", "Csv filename")
	flag.Parse()

	if len(*csvFile) <= 0 {
		fmt.Printf("Invalid file. got=%v\n", *csvFile)
		os.Exit(1)
	}

	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		fmt.Printf("Unknown output type. got)%v\n", *outputType)
	}
	if err != nil {
		fmt.Printf("Could not create generator: %v\n", err)
		os.Exit(1)
	}

	certs, err := cert.ParseCSV(*csvFile)
	if err != nil {
		fmt.Printf("Error during parsing csv: %v\n", err)
		os.Exit(1)
	}
	for _, cert := range certs {
		err = saver.Save(*cert)
		if err != nil {
			fmt.Printf("Error during saving files: %v\n", err)
			os.Exit(1)
		}
	}
}
