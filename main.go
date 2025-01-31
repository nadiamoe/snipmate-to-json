package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

type snippet struct {
	Name   string `json:"-"`
	Prefix string `json:"prefix"`
	Body   string `json:"body"`
}

func run() error {
	scanner := bufio.NewScanner(os.Stdin)

	snippets := map[string]snippet{}
	var currentSnippet snippet

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \n")

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "snippet") {
			if currentSnippet.Prefix != "" {
				currentSnippet.Body = strings.TrimSpace(currentSnippet.Body)
				snippets[currentSnippet.Name] = currentSnippet
			}

			spec := strings.Fields(line)
			currentSnippet = snippet{
				Prefix: spec[1],
				Name:   strings.Trim(strings.Join(spec[2:], " "), "\""),
			}

			continue
		}

		currentSnippet.Body += strings.TrimPrefix(line, "\t") + "\n"
	}

	currentSnippet.Body = strings.TrimSpace(currentSnippet.Body)
	snippets[currentSnippet.Name] = currentSnippet

	return json.NewEncoder(os.Stdout).Encode(snippets)
}
