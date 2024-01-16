package main

import (
    "testing"
	"strings"
)

func TestFileContentParser(t *testing.T) {
    got := parseFile(strings.Split("it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.", " "))
	want := "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = parseFile(strings.Split("Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", " "))
	want = "Simply add 66 and 2 and you will see the result is 68."

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = parseFile(strings.Split("There is no greater agony than bearing a untold story inside you.", " "))
	want = "There is no greater agony than bearing an untold story inside you."

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = parseFile(strings.Split("Punctuation tests are ... kinda boring ,don't you think !?", " "))
	want = "Punctuation tests are... kinda boring, don't you think!?"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}