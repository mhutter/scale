package main

import (
	"strings"
	"testing"
)

func TestBuildSequenceCmajor(t *testing.T) {
	expected := "CDEFGABC"
	actual := strings.Join(buildSequence("C", "major"), "")
	if actual != expected {
		t.Errorf("\nexpected - %s\n  actual - %s", expected, actual)
	}
}

func TestBuildSequenceGminor(t *testing.T) {
	expected := "GAA#CDD#FG"
	actual := strings.Join(buildSequence("G", "minor"), "")
	if actual != expected {
		t.Errorf("\nexpected - %s\n  actual - %s", expected, actual)
	}
}
