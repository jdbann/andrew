package content

import (
	"context"
	"time"

	"encore.app/content/queries"
)

// CreatePostParams is the request data for the CreatePost endpoint.
type CreatePostParams struct {
	// Unique identifier for the post.
	Slug string `json:"slug"`
	// Main title for the post.
	Title string `json:"title"`
	// One-paragraph summary of the post's content.
	Summary string `json:"summary"`
	// Full content of the post.
	Body string `json:"body"`
}

// CreatePostResponse is the response data for the CreatePost endpoint.
type CreatePostResponse struct {
	// Unique identifier for the post.
	Slug string `json:"slug"`
	// Main title for the post.
	Title string `json:"title"`
	// One-paragraph summary of the post's content.
	Summary string `json:"summary"`
	// Full content of the post.
	Body string `json:"body"`
	// When the post was created.
	CreatedAt time.Time `json:"created_at"`
}

// CreatePost adds a new post to the database.
//
//encore:api public method=POST path=/content/posts
func CreatePost(ctx context.Context, params *CreatePostParams) (*CreatePostResponse, error) {
	post, err := queries.New(contentDB.Stdlib()).CreatePost(ctx, queries.CreatePostParams{
		Slug:    params.Slug,
		Title:   params.Title,
		Summary: params.Summary,
		Body:    params.Body,
	})
	if err != nil {
		return nil, err
	}

	return &CreatePostResponse{
		Slug:      post.Slug,
		Title:     post.Title,
		Summary:   post.Summary,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
	}, nil
}
