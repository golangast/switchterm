package gentil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//f is for file, o is for old text, n is for new text
func UpdateText(f string, o string, n string) error {
	input, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}

	output := bytes.Replace(input, []byte(o), []byte(n), -1)

	if err = ioutil.WriteFile(f, output, 0666); err != nil {
		fmt.Println(err)
	}

	return nil
}
func FindTextNReturn(p, str string) string {
	// Open file for reading.
	var file, err = os.OpenFile(p, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	toplevel := TrimDot(str)
	property := TrimDotright(str)
	strs := strings.Replace(property, ".", " ", 1)
	// fmt.Println(str)
	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		if strings.Contains(string(text), toplevel) {
			//is the dot string and split it
			if strings.Contains(string(text), strs) {
				return string(text)
			}
		}
		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			fmt.Println(err)

		}
	}

	// fmt.Println("Reading from file.")
	fmt.Println(string(text))

	return ""
}
func TrimDot(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}
func TrimDotright(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[idx:]
	}
	return s
}
