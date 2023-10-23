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
	fs := fstest.MapFS{
		"post-1.md": {Data: []byte("Title: Post 1")},
		"post-2.md": {Data: []byte("Title: Post 2")},
	}
	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := blogposts.Post{Title: "Post 1"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}