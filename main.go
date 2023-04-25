package main

import (
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var saltSuffix = ""
	var saltPrefix = ""
	var dirty bool = false
	var listSsnHash bool = false

	flag.BoolVar(&dirty, "d", false, "dirty mode, display cleartext in output")
	flag.BoolVar(&listSsnHash, "l", false, "list ssn + hash equivalent in ssn_hash.txt")
	flag.StringVar(&saltPrefix, "p", "", "saltstring prefix")
	flag.StringVar(&saltSuffix, "s", "", "saltstring suffix")

	flag.Parse()

	if saltPrefix == "" && saltSuffix == "" {
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
		hash := hashData(saltPrefix, s, saltSuffix)
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
			fmt.Printf("tried to create ssn_hash.txt for output of ssn hash file, but got error: %v\nUsing stdout.", err)
			ssnHashOut = os.Stdout
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

func hashData(saltPrefix string, s string, saltSuffix string) string {
	h := sha256.New()
	h.Write([]byte(saltPrefix + s + saltSuffix))
	fmt.Printf("%v%v%v\n", saltPrefix, s, saltSuffix)
	return fmt.Sprintf("%x", h.Sum(nil))
}
