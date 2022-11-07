package content_test

import (
	"context"
	"testing"
	"time"

	"encore.app/content"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreatePost(t *testing.T) {
	setupDB(t)

	res, err := content.CreatePost(context.Background(), &content.CreatePostParams{
		Slug:    "a-test-post",
		Title:   "A test post",
		Summary: "Just verifying that creating a post works.",
		Body:    "Sometimes it's worth being careful. Adding tests is a great way to do that.\n\nEspecially if they're run automatically!",
	})
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(&content.CreatePostResponse{
		Post: content.Post{
			Slug:      "a-test-post",
			Title:     "A test post",
			Summary:   "Just verifying that creating a post works.",
			Body:      "Sometimes it's worth being careful. Adding tests is a great way to do that.\n\nEspecially if they're run automatically!",
			CreatedAt: time.Now(),
		},
	}, res, cmpopts.EquateApproxTime(time.Second)); diff != "" {
		t.Error(diff)
	}
}

func TestGetPosts(t *testing.T) {
	setupDB(t)

	_, err := content.CreatePost(context.Background(), &content.CreatePostParams{
		Slug:    "a-test-post",
		Title:   "A test post",
		Summary: "Just verifying that creating a post works.",
		Body:    "Sometimes it's worth being careful. Adding tests is a great way to do that.\n\nEspecially if they're run automatically!",
	})
	if err != nil {
		t.Fatal(err)
	}

	res, err := content.GetPosts(context.Background())

	if diff := cmp.Diff(&content.GetPostsResponse{
		Posts: []content.Post{
			{
				Slug:      "a-test-post",
				Title:     "A test post",
				Summary:   "Just verifying that creating a post works.",
				Body:      "Sometimes it's worth being careful. Adding tests is a great way to do that.\n\nEspecially if they're run automatically!",
				CreatedAt: time.Now(),
			},
		},
	}, res, cmpopts.EquateApproxTime(time.Second)); diff != "" {
		t.Error(diff)
	}
}
