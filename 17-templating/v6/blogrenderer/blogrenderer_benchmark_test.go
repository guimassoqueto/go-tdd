package blogrenderer_test

import (
	"blogrender/blogrenderer"
	"io"
	"testing"
)

func BenchmarkRender(b *testing.B) {
	var (
		post = blogrenderer.Post{
			Title: "Any Title",
			Body: "Any Body",
			Tags: []string{"tdd", "golang"},
		}
	)

	blogRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blogRenderer.Render(io.Discard, post)
	}

}