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