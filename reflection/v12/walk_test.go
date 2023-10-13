package main

import (
	"reflect"
	"slices"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}


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
		{
			Description: "nested fields",
			Input: Person{
				Name: "Chris",
				Profile: Profile{
					Age: 33,
					City: "London",
				},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Description: "pointer to things",
			Input: &Person{
				Name: "Chris",
				Profile: Profile{
					Age: 33,
					City: "London",
				},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Description: "slices",
			Input: []Profile{
				{Age: 33, City: "London"},
				{Age: 44, City: "Rio de Janeiro"},
			},
			ExpectedCalls: []string{"London", "Rio de Janeiro"},
		},
		{
			Description: "arrays",
			Input: [2]Profile{
				{Age: 33, City: "London"},
				{Age: 44, City: "Rio de Janeiro"},
			},
			ExpectedCalls: []string{"London", "Rio de Janeiro"},
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

		t.Run("with maps", func(t *testing.T) {
			aMap := map[string]string{
				"Cow": "Moo",
				"Sheep": "Bee",
			}
			var got []string
			Walk(aMap, func(input string){
				got = append(got, input)
			})
			assertContains(t, got, "Moo")
			assertContains(t, got, "Bee")
		})

		t.Run("with channels", func(t *testing.T) {
			aChannel := make(chan Profile)
			go func() {
				aChannel <- Profile{Age: 33, City: "Berlin"}
				aChannel <- Profile{Age: 34, City: "New York"}
				close(aChannel)
			}()
			var got []string
			want := []string{"Berlin", "New York"}

			Walk(aChannel, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("with functions", func(t *testing.T) {
			aFunction := func() (Profile, Profile) {
				return Profile{Age: 33, City: "Berlin"}, Profile{Age: 18, City: "New York"}
			}

			var got []string
			want := []string{"Berlin", "New York"}

			Walk(aFunction, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func assertContains(t testing.TB, got []string, want string) {
	t.Helper()
	contains := slices.Contains(got, want)
	if !contains {
		t.Errorf("expect %v contains %v, but it doesn't", got, want)
	}
}