package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 || !strings.HasSuffix(os.Args[1], ".pls") {
		fmt.Printf("usage: %s <file.pls>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if rawbytes, err := os.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(rawbytes))
	}
}
