package md

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"os"
)

func ProcessFile(relPath string) string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("could not get pwd")
		return ""
	}

	source, err := os.ReadFile(pwd + relPath)
	if err != nil {
		fmt.Printf("could not read file: [%v]\n", pwd+relPath)
		return ""
	}
	return string(parseMd(source))
}

func parseMd(source []byte) []byte {

	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
		panic(err)
	}

	return buf.Bytes()
}
