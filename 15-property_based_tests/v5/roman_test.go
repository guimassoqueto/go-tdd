package main 

import "testing"

var testCases = []struct{
	description string
	arabic int
	roman string
	}{
	{ description: "1 gets converted to I", arabic: 1, roman: "I"},
	{ description: "2 gets converted to II", arabic: 2,	roman: "II"},
	{ description: "3 gets converted to III",	arabic: 3, roman: "III"},
	{ description: "4 gets converted to IV",	arabic: 4, roman: "IV"},
	{ description: "5 gets converted to V",	arabic: 5, roman: "V"},
	{ description: "8 gets converted to VIII",	arabic: 8, roman: "VIII"},
	{ description: "9 gets converted to IX",	arabic: 9, roman: "IX"},
	{ description: "40 gets converted to XL",	arabic: 40, roman: "XL"},
	{ description: "47 gets converted to XLVII",	arabic: 47, roman: "XLVII"},
	{ description: "49 gets converted to XLIX",	arabic: 49, roman: "XLIX"},
	{ description: "50 gets converted to L",	arabic: 50, roman: "L"},
	{ description: "100 gets converted to C",	arabic: 100, roman: "C"},
	{ description: "90 gets converted to XC",	arabic: 90, roman: "XC"},
	{ description: "400 gets converted to CD",	arabic: 400, roman: "CD"},
	{ description: "500 gets converted to D",	arabic: 500, roman: "D"},
	{ description: "900 gets converted to CM",	arabic: 900, roman: "CM"},
	{ description: "1000 gets converted to M",	arabic: 1000, roman: "M"},
	{ description: "1984 gets converted to MCMLXXXIV",	arabic: 1984, roman: "MCMLXXXIV"},
	{ description: "2014 gets converted to MMXIV",	arabic: 2014, roman: "MMXIV"},
	{ description: "1006 gets converted to MVI",	arabic: 1006, roman: "MVI"},
	{ description: "798 gets converted to DCCXCVIII",	arabic: 798, roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
		for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			got := ConvertToRoman(testCase.arabic)
			want := testCase.roman
			
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}

func TestArabicNumerals(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			got := ConvertToArabic(tt.roman)
			want := tt.arabic

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}
