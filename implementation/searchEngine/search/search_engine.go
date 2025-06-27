package search

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Engine struct {
	dirpath      string
	SearchResult map[string]string
}

func InitializeEngine(dirpath string) *Engine {
	return &Engine{
		dirpath:      dirpath,
		SearchResult: make(map[string]string),
	}
}

func (e *Engine) Search(query string) map[string]string {

	dir, err := os.ReadDir(e.dirpath)

	if err != nil {
		log.Fatal(err)
	}

	searchResult := make(map[string]string)

	for _, v := range dir {
		info, _ := v.Info()
		filename := info.Name()
		counter := 0
		if _, isTxt := strings.CutSuffix(filename, ".txt"); isTxt {
			search := func(filename string) {
				path := filepath.Join("search", "data", filename)
				file, err := os.Open(path)
				if err != nil {
					log.Fatal(err)
					log.Fatal("File Not Found")
				}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					line := scanner.Text()
					counter++
					// if found := strings.Contains(line, query); found {
					// 	key := fmt.Sprintf("file: %s, query: %s, line: %d", filename, query, counter)
					// 	searchResult[key] = line
					// }

					if position := Find(query, line); position != -1 {
						key := fmt.Sprintf("file:- %s, query:- %s, (line, index):- (%d, %d)", filename, query, counter, position)
						searchResult[key] = line
					}
				}
			}
			search(filename)
		}
	}
	return searchResult
}

func Find(pattern, line string) int {
	pattern_len := len(pattern) // Get the length of the pattern to search for
	p_mask := make([]int, 400)  // Create a mask array for all possible characters (ASCII range)
	// bit 0 marks that we've matched 0 characters so far
	A := ^1 // Initialize bitmask A with all bits set except the lowest bit (bit 0)

	if pattern_len == 0 {
		return -1 // If the pattern is empty, return -1 (not found)
	}

	if pattern_len > 63 {
		log.Fatal("Pattern to Search is too long") // Limit pattern length to 63 for bitmasking
	}

	for i := 0; i <= 299; i++ {
		p_mask[i] = ^0 // Set all bits to 1 for each character in the mask
	}

	for i := 0; i < pattern_len; i++ {
		p_mask[int(pattern[i])] &= ^(1 << i) // Update mask for each character in the pattern
	}

	for i := 0; i < len(line); i++ {
		A |= p_mask[int(line[i])]    // Update A with the mask for the current character
		A <<= 1                      // Shift A left by 1 bit (progressing the match)
		if A&(1<<pattern_len) == 0 { // Check if the pattern has been matched
			return i - pattern_len + 1 // Return the starting index of the match
		}
	}
	return -1 // Return -1 if the pattern is not found
}
