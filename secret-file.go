package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

//go:embed secrets/secret-key.txt
//go:embed secrets/*
var secretFS embed.FS

const uncovered = "uncovered"

func UnearthFiles() {
	fs.WalkDir(secretFS, "secrets", func(path string, d fs.DirEntry, err error) error {
		os.Mkdir(uncovered, 0777)
		switch path {
		// skip root and secret-key.txt, we don't want to copy those
		case "secrets":
		case "secrets/secret-key.txt":
			return nil
		default:
			var file, _ = secretFS.ReadFile(path)
			var filename = strings.Split(path, "/")[1]

			os.WriteFile(uncovered+"/"+filename, file, 0777)
		}

		return nil
	})
}

func main() {

	secretPtr := flag.String("secret", "", "secret to uncover files (required)")

	// if no args are passed, exit
	if len(os.Args) < 2 {
		fmt.Println("No args passed. Expected 'secret'")
		os.Exit(1)
	}

	// parse the flags
	flag.Parse()
	// fetch the secret
	var secret, _ = secretFS.ReadFile("secrets/secret-key.txt")
	// if the secret is correct, uncover the files, else exit
	if string(secret) == *secretPtr {
		UnearthFiles()
		fmt.Println("Files uncovered, look for the 'uncovered' directory")
		os.Exit(0)
	} else {
		fmt.Println("Failed to uncover files. Incorrect secret")
		os.Exit(1)
	}
}
