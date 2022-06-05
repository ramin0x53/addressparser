package parse

import (
	"reflect"
	"testing"
)

func TestAddrParse(t *testing.T) {
	got := AddrParse("The Book Club 100-106 Leonard St Shoreditch London EC2A 4RH, United Kingdom")

	want := map[string]string{
		"city":         "london",
		"country":      "united kingdom",
		"house":        "the book club",
		"house_number": "100-106",
		"postcode":     "ec2a 4rh",
		"road":         "leonard st",
		"suburb":       "shoreditch",
	}
	if !(reflect.DeepEqual(got, want)) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
