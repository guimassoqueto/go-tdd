package main

import "errors"

type Dict map[string]string

func (dict Dict) Search(word string) (string, error) {
	definition, ok := dict[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}

