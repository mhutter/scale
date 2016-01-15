package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	// the chromatic scale
	notes = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

	// scale steps
	steps = map[string][]int{
		"major": []int{2, 2, 1, 2, 2, 2, 1},
		"minor": []int{2, 1, 2, 2, 1, 2, 2},
	}

	// known scales - see init()
	scales []string

	// command line arguments
	root  string
	scale string
)

// Get and store command line arguments
func init() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	for scale, _ := range steps {
		scales = append(scales, scale)
	}

	scale = strings.ToLower(os.Args[1])
	if !has(scales, scale) {
		fmt.Printf("Unknown Scale '%s'.\n", scale)
		usage()
		os.Exit(2)
	}

	root = strings.ToUpper(os.Args[2])
	if !has(notes, root) {
		fmt.Printf("Unknown Root Note '%s'.\n", root)
		usage()
		os.Exit(3)
	}
}

func main() {
	fmt.Printf("%s %s scale: ", root, scale)

	// prepare the output
	length := len(steps[scale]) + 1
	out := make([]string, length)

	// find the index of the root note
	var index int
	for i, e := range notes {
		if e == root {
			index = i
			break
		}
	}

	for i, n := range steps[scale] {
		out[i] = notes[index]
		index = (index + n) % len(notes)
	}
	out[length-1] = notes[index]
	fmt.Println(strings.Join(out, " "))
}

// Print usage information
func usage() {
	fmt.Printf("Usage: %s SCALE ROOT\n\n", os.Args[0])
	fmt.Printf("    Scales: %s\n", strings.Join(scales, ", "))
	fmt.Printf("     Notes: %s\n", strings.Join(notes, ", "))
}

// Check if `needle` is in `haystack`.
func has(haystack []string, needle string) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}
	return false
}
