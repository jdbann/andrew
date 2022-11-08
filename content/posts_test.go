package content_test

import (
	"context"
	"testing"
	"time"

	"encore.app/content"
	"encore.dev/beta/errs"
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

func TestGetPost(t *testing.T) {
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

	tests := []struct {
		name        string
		slug        string
		wantRes     *content.GetPostResponse
		wantErrCode errs.ErrCode
	}{
		{
			name: "success",
			slug: "a-test-post",
			wantRes: &content.GetPostResponse{
				Post: content.Post{
					Slug:      "a-test-post",
					Title:     "A test post",
					Summary:   "Just verifying that creating a post works.",
					Body:      "Sometimes it's worth being careful. Adding tests is a great way to do that.\n\nEspecially if they're run automatically!",
					CreatedAt: time.Now(),
				},
			},
			wantErrCode: errs.OK,
		},
		{
			name:        "not found",
			slug:        "unknown-post",
			wantRes:     nil,
			wantErrCode: errs.NotFound,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			res, err := content.GetPost(context.Background(), tt.slug)

			if tt.wantErrCode != errs.Code(err) {
				t.Errorf("want error code %q; got %q", tt.wantErrCode.String(), errs.Code(err).String())
			}

			if diff := cmp.Diff(tt.wantRes, res, cmpopts.EquateApproxTime(time.Second)); diff != "" {
				t.Error(diff)
			}
		})
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
