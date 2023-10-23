package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dict{"test": "this is just a test"}	

	t.Run("should return the expected key", func(t *testing.T) {
		assertKeyValue(t, dict, "test", "this is just a test")
	})

	t.Run("unknown key", func(t *testing.T) {
		_ , err := dict.Search("not-valid")
		want := ErrNotFound
		
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		dict := Dict{}	
		err := dict.Add("key", "value")

		assertKeyValue(t, dict, "key", "value")
		assertNil(t, err)
	})

	t.Run("trying to add to an existing key should return an error", func(t *testing.T) {
		key := "test-key"
		value := "any-value"
		dict := Dict{key: value}	
		err := dict.Add(key, "new-value")

		assertKeyValue(t, dict, key, value)
		assertError(t, err, ErrKeyAlreadyExists)
	})
}

func assertKeyValue(t testing.TB, dict Dict, key, value string) {
	t.Helper()
	got, err := dict.Search(key)

	if got != value {
		t.Errorf("got %q want %q", got, value)
	}

	if err != nil {
		t.Errorf("got %q want nil", err)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNil(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got %q want nil", err)
	}
}