package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	ignore "github.com/sabhiram/go-gitignore"
)

const marker = "##DEV_TREE_OUTPUT##"

func isBinary(data []byte) bool {
	for _, b := range data {
		if b == 0 {
			return true
		}
	}
	return false
}

func isBinaryFile(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	buf := make([]byte, 8000)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return false, err
	}
	return isBinary(buf[:n]), nil
}

func main() {
	outputFile := flag.String("output", "structure.txt", "Output file for directory structure")
	flag.Parse()

	ign, err := ignore.CompileIgnoreFile(".gitignore")
	if err != nil {
		log.Println("No .gitignore found; proceeding without ignore rules")
	}

	outAbs, err := filepath.Abs(*outputFile)
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}

	out, err := os.Create(*outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer out.Close()

	// Write identification marker at the top.
	fmt.Fprintln(out, marker)

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the output file by comparing absolute paths.
		currAbs, err := filepath.Abs(path)
		if err == nil && currAbs == outAbs {
			return nil
		}

		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		if ign != nil && ign.MatchesPath(path) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		rel, _ := filepath.Rel(".", path)
		if info.IsDir() {
			fmt.Fprintf(out, "[%s]\n", info.Name())
		} else {
			if info.Name() == ".gitignore" {
				return nil
			}
			fmt.Fprintf(out, "- %s\n", info.Name())

			isBin, err := isBinaryFile(path)
			if err != nil {
				fmt.Fprintf(out, "Error checking file type: %v\n", err)
				return nil
			}
			if isBin {
				fmt.Fprintf(out, "----- SKIPPED BINARY FILE: %s -----\n", rel)
				return nil
			}

			fmt.Fprintf(out, "----- START OF FILE: %s -----\n", rel)
			content, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Fprintf(out, "Error reading file: %v\n", err)
			} else {
				fmt.Fprintf(out, "%s\n", string(content))
			}
			fmt.Fprintf(out, "----- END OF FILE: %s -----\n", rel)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking directory: %v", err)
	}
}
