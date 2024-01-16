// run tests using 'go test *.go'
package main

import (
    "testing"
	"main/numConverter"
	"main/wordConverter"
)

func TestConvertFromBinary(t *testing.T) {
    got := numConverter.ConvertFromBinary("1")
	want := "1"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = numConverter.ConvertFromBinary("10")
	want = "2"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = numConverter.ConvertFromBinary("101")
	want = "5"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
	got = numConverter.ConvertFromBinary("101011")
	want = "43"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestConvertFromHex(t *testing.T) {
    got := numConverter.ConvertFromHex("21C")
	want := "540"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = numConverter.ConvertFromHex("1E")
	want = "30"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = numConverter.ConvertFromHex("42")
	want = "66"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCapitalizationChanger(t *testing.T) {
    got := wordConverter.CaseChanger("up", "go")
	want := "GO"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = wordConverter.CaseChanger("low", "SHOUTING")
	want = "shouting"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = wordConverter.CaseChanger("cap", "bridge")
	want = "Bridge"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCheckIfaORan(t *testing.T) {
    got := wordConverter.CheckIfaORan("amazing")
	want := "an"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = wordConverter.CheckIfaORan("Amazing")
	want = "an"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = wordConverter.CheckIfaORan("house")
	want = "a"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}