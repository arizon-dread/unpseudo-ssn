package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func main() {
	var salt = ""
	var dirty bool = false
	var listSsnHash bool = false
	if len(os.Args) > 1 {

		for _, arg := range os.Args {
			if strings.Contains(arg, "-d") {
				dirty = true
			} else if strings.Contains(arg, "-l") {
				listSsnHash = true
			} else {
				salt = arg
			}
		}
		//fmt.Printf("dirty: %v, list: %v, salt: %v", dirty, listSsnHash, salt)

	}
	if salt == "" {
		fmt.Println("Running hash without salt.")
	}
	fi, err := os.Open("input.txt")
	if err != nil {
		panic("Could not read input.txt")
	}
	defer fi.Close()

	fo, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf("Could not create output.txt: %v, using stdout.\n", err)
		fo = os.Stdout
	}

	defer fo.Close()

	data, err := os.Open("hashed_data.txt")
	if err != nil {
		fmt.Printf("Could not read hashed_data.txt: %v\n", err)
		panic("Could not read hashed_data.txt")
	}

	defer data.Close()

	ssn := readFileToStringSlice(fi)

	var ssnHash = make(map[string]string)
	for _, s := range ssn {
		hash := hashData(s, salt)
		ssnHash[hash] = s
	}

	content := readFileToStringSlice(data)

	for _, line := range content {
		//fmt.Printf("line %v\n", line)
		for key, val := range ssnHash {
			// fmt.Printf("key %v\n", key)
			// fmt.Printf("val %v\n", val)
			if strings.Contains(line, key) {
				var newline string = ""
				if dirty {
					newline = strings.Replace(line, key, val, 1)
				} else {
					newline = line
				}

				// fmt.Printf("line w/ replaced hash: %v\n", newline)
				fo.WriteString(newline + "\n")
			}
		}
	}
	if listSsnHash {
		ssnHashOut, err := os.Create("ssn_hash.txt")
		if err != nil {
			fmt.Printf("tried to create ssn_hash.txt for output of ssn hash file, but got error: %v\n", err)
		}

		defer ssnHashOut.Close()
		for key, val := range ssnHash {
			ssnHashOut.WriteString(fmt.Sprintf("%v %v\n", val, key))
		}
	}
}

func readFileToStringSlice(f *os.File) []string {
	var content []string
	datascanner := bufio.NewScanner(f)
	for datascanner.Scan() {
		t := datascanner.Text()
		content = append(content, t)
	}
	return content
}

func hashData(s string, salt string) string {
	h := sha256.New()
	h.Write([]byte(s + salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}
