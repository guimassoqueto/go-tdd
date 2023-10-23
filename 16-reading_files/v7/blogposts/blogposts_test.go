package blogposts_test

import (
	"errors"
	blogposts "grf/blogposts"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("always fail")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = "Title: Post1\nDescription: Description1\nTags: golang, web\n---\nHello\nWorld"
		secondBody = "Title: Post2\nDescription: Description2\nTags: tdd, test\n---\nB\nL\nN"
	)

	fs := fstest.MapFS{
		"post-1.md": {Data: []byte(firstBody)},
		"post-2.md": {Data: []byte(secondBody)},
	}
	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := blogposts.Post{
		Title: "Post1", 
		Description: "Description1", 
		Tags: []string{"golang", "web"},
		Body: "Hello\nWorld",
	}
	assertPost(t, got, want)
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}