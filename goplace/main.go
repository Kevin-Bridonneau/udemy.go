/*
--------------------------------------------------
	First project in go Goplace (find and replace)
	by Kevin Bridonneau 30/12/2020
--------------------------------------------------
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var fileName string = "testFile.txt"
	var old string = "Python"
	var new string = "Go"

	occ, lines, err := findReplaceFile(fileName, old, new)
	check(err)

	fmt.Println("== Summary ==")
	fmt.Printf("Number of occurrences of Python: %v\n", occ)
	fmt.Printf("Number of lines: %v\n", len(lines))
	fmt.Printf("Lines: %v\n", lines)
	fmt.Println("== End of Summary ==")
}

func findReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	var lineNbr int
	// open the files and prepare to close at the end of this func
	f, err := os.Open(src)
	defer f.Close()
	//handle errors while opening
	check(err)

	nf, err := os.OpenFile("newfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer nf.Close()
	//handle errors while opening
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		found, res, tempOcc := processLine(line, old, new)
		if found == true {
			lines = append(lines, lineNbr)
			line = res
			occ += tempOcc
		}
		_, err = nf.WriteString(line)
		_, err = nf.WriteString("\n")
		check(err)

		lineNbr++
	}
	return occ, lines, err
}

func processLine(line, old, new string) (found bool, res string, occ int) {
	found = strings.Contains(line, old)
	if found == false {
		return found, line, occ
	}
	occ += strings.Count(line, old)
	res = strings.Replace(line, old, new, -1 )
	return found, res, occ
}
