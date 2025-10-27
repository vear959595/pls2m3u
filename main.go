package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Song struct {
	Title    string
	Filename string
	Seconds  int
}

func main() {
	if len(os.Args) == 1 || !strings.HasSuffix(os.Args[1], ".pls") {
		fmt.Printf("usage: %s <file.pls>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if rawbytes, err := os.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", string(rawbytes))
	}
}

func readPlsPlaylist(data string) (songs []Song) {
	var song Song
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "[playlist]") {
			continue
		}
	}
}
