package main

import "testing"

var cases = []struct {
	Description string
	Arabic      int
	Roman       string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestConvertArabicToRomanNumerals(t *testing.T) {
	t.Logf("Test Convert Arabic To Roman Numerals")

	for _, test := range cases {
		t.Logf("\t%s", test.Description)

		got := ConvertToRoman(test.Arabic)
		want := test.Roman

		if got == want {
			t.Logf("\t%s\tGot %s", succed, got)
		} else {
			t.Errorf("\t%s\tGot %s, want %s", failed, got, want)
		}
	}
}

func TestConvertRomanToArabicNumerals(t *testing.T) {
	t.Logf("Test Convert Roman To Arabic Numerals")

	for _, test := range cases[:5] {
		t.Logf("\t%s", test.Description)

		got := ConvertToArabic(test.Roman)
		want := test.Arabic

		if got == want {
			t.Logf("\t%s\tGot %d", succed, got)
		} else {
			t.Errorf("\t%s\tGot %d, want %d", failed, got, want)
		}
	}
}
