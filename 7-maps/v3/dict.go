package main

import "errors"

type Dict map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (dict Dict) Search(word string) (string, error) {
	definition, ok := dict[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

