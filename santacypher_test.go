package santacypher

import (
	"testing"
)

func TestToHoHoHo(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "Hello Gopher", want: "HoHoHoHo🎅Ho🎅HoHOHoHo🎅HoHOHoHo🎅HOHOHO🎅 HOHOHo🎅HOHOHO🎅HoHOHOHo🎅HoHoHoHo🎅Ho🎅HoHOHo🎅"},
		{input: "HELLO GOPHER", want: "HoHoHoHo🎅Ho🎅HoHOHoHo🎅HoHOHoHo🎅HOHOHO🎅 HOHOHo🎅HOHOHO🎅HoHOHOHo🎅HoHoHoHo🎅Ho🎅HoHOHo🎅"},
		{input: "🎅", want: "🎅"},
	}
	for _, test := range tests {
		got := ToHoHoHo(test.input)
		if got != test.want {
			t.Fatalf("want %v, but %v:", test.want, got)
		}
	}
}

func TestFromHoHoHo(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "HoHoHoHo🎅Ho🎅HoHOHoHo🎅HoHOHoHo🎅HOHOHO🎅 HOHOHo🎅HOHOHO🎅HoHOHOHo🎅HoHoHoHo🎅Ho🎅HoHOHo🎅", want: "HELLO GOPHER"},
		{input: "HoHoHoHo🎅Ho🎅HoHOHoHo🎅HoHOHoHo🎅HOHOHO🎅 HOHOHo🎅HOHOHO🎅HoHOHOHo🎅HoHoHoHo🎅Ho🎅HoHOHo🎅", want: "HELLO GOPHER"},
		{input: "🎅", want: ""},
	}
	for _, test := range tests {
		got := FromHoHoHo(test.input)
		if got != test.want {
			t.Fatalf("want %v, but %v:", test.want, got)
		}
	}
}
