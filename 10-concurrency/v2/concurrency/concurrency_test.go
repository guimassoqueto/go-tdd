package concurrency

import (
	"reflect"
	"testing"
)

func mockCheckWebsite(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://yahoo.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"https://google.com": true,
		"https://yahoo.com": true,
		"waat://furhurterwe.geds": false,
	}

	got := CheckWebsite(mockCheckWebsite, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
