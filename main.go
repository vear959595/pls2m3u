package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 || !strings.HasPrefix(os.Args[1], ".pls") {
		fmt.Printf("usage: %s <file.pls>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
}
