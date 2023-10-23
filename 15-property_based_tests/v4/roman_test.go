package main 

import "testing"

func TestRomanNumerals(t *testing.T) {
	testCases := []struct{
		description string
		arabic int
		want string
	}{
		{ description: "1 gets converted to I", arabic: 1, want: "I"},
		{ description: "2 gets converted to II", arabic: 2,	want: "II"},
		{ description: "3 gets converted to III",	arabic: 3, want: "III"},
		{ description: "4 gets converted to IV",	arabic: 4, want: "IV"},
		{ description: "5 gets converted to V",	arabic: 5, want: "V"},
		{ description: "8 gets converted to VIII",	arabic: 8, want: "VIII"},
		{ description: "9 gets converted to IX",	arabic: 9, want: "IX"},
		{ description: "40 gets converted to XL",	arabic: 40, want: "XL"},
		{ description: "47 gets converted to XLVII",	arabic: 47, want: "XLVII"},
		{ description: "49 gets converted to XLIX",	arabic: 49, want: "XLIX"},
		{ description: "50 gets converted to L",	arabic: 50, want: "L"},
		{ description: "100 gets converted to C",	arabic: 100, want: "C"},
		{ description: "90 gets converted to XC",	arabic: 90, want: "XC"},
		{ description: "400 gets converted to CD",	arabic: 400, want: "CD"},
		{ description: "500 gets converted to D",	arabic: 500, want: "D"},
		{ description: "900 gets converted to CM",	arabic: 900, want: "CM"},
		{ description: "1000 gets converted to M",	arabic: 1000, want: "M"},
		{ description: "1984 gets converted to MCMLXXXIV",	arabic: 1984, want: "MCMLXXXIV"},
		{ description: "2014 gets converted to MMXIV",	arabic: 2014, want: "MMXIV"},
		{ description: "1006 gets converted to MVI",	arabic: 1006, want: "MVI"},
		{ description: "798 gets converted to DCCXCVIII",	arabic: 798, want: "DCCXCVIII"},

	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := ConvertToRoman(testCase.arabic)
			want := testCase.want
			assertEqual(t, got, want)
		})
	}
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}