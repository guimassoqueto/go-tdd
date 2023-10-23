package blogposts_test

import (
	blogposts "grf/blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world-1.md": {Data: []byte("Hi")},
		"hello-world-2.md": {Data: []byte("Hello")},
	}
	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted, %d posts", len(posts), len(fs))
	}
}