package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	testCases := []struct {
		Description string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			Description: "struct with one field",
			Input: struct{
				Name string
			}{ Name: "Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Description: "struct with two string fields",
			Input: struct{
				Name string
				City string
			}{ Name: "Chris", City: "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Description: "struct with no string field",
			Input: struct{
				Name string
				Age int
			}{ Name: "Chris", Age: 35},
			ExpectedCalls: []string{"Chris"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			var got []string
			Walk(testCase.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, testCase.ExpectedCalls) {
				t.Errorf("got %v want %v", got, testCase.ExpectedCalls)
			}
		})
	}
}