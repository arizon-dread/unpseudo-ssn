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
	if len(os.Args) > 1 {
		salt = os.Args[1]
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
		fmt.Printf("Could not create output.txt\n")
	}

	defer fo.Close()

	data, err := os.Open("hashed_data.txt")
	if err != nil {
		fmt.Printf("Could not read hashed_data.txt\n")
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
				newline := strings.Replace(line, key, val, 1)
				// fmt.Printf("line w/ replaced hash: %v\n", newline)
				fo.WriteString(newline + "\n")
			}
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
