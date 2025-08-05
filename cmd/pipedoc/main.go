package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"pipedoc/internal/generator"
	"pipedoc/internal/parser"
)

func main() {
	input := flag.String("f", ".github/workflows/ci.yml", "Path to CI file")
	output := flag.String("o", "docs/ci-pipeline.md", "Output documentation path")
	flag.Parse()

	pipeline, err := parser.ParseGitHubActions(*input)
	if err != nil {
		log.Fatalf("Error parsing CI file: %v", err)
	}

	md, mermaid := generator.GenerateMarkdown(pipeline)
	os.MkdirAll("docs", os.ModePerm)

	if err := os.WriteFile(*output, []byte(md), 0644); err != nil {
		log.Fatalf("Error writing markdown: %v", err)
	}

	if err := os.WriteFile("docs/ci-pipeline.mmd", []byte(mermaid), 0644); err != nil {
		log.Fatalf("Error writing mermaid file: %v", err)
	}

	fmt.Println("CI documentation generated in docs/")
}
