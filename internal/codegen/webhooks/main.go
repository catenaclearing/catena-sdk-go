package main

import (
	"flag"
	"log"
	"os"
	"sort"
)

func main() {
	specPath := flag.String("spec", "", "Path to OpenAPI spec")
	outPath := flag.String("out", "", "Path to output file")
	flag.Parse()

	if *specPath == "" || *outPath == "" {
		log.Fatal("Usage: generator --spec <path> --out <path>")
	}

	events, err := discoverWebhookEvents(*specPath)
	if err != nil {
		log.Fatalf("Failed to discover events: %v", err)
	}

	// Sort for deterministic output
	sort.Slice(events, func(i, j int) bool {
		return events[i].EventName < events[j].EventName
	})

	code, err := renderEvents(events)
	if err != nil {
		log.Fatalf("Failed to render code: %v", err)
	}

	if err := os.WriteFile(*outPath, code, 0644); err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}
