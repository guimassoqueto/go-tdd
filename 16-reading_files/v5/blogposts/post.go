package blogposts

import (
	"bufio"
	"io"
)

type Post struct{
	Title string
	Description string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	scanner.Scan()
	title := scanner.Text()

	scanner.Scan()
	description := scanner.Text()

	return Post{Title: title[7:], Description: description[13:]}, nil
}