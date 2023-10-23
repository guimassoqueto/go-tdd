package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct{
	Title string
	Description string
	Tags []string
	Body string
}

const (
	titleSeparator = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator = "Tags: "
	bodySeparator = "Body: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")
	body := ReadBody(scanner)

	return Post{
		Title: title, 
		Description: description,
		Tags: tags,
		Body: body,
		}, nil
}

func ReadBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore line with ---
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}