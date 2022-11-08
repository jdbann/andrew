package content

import (
	"context"
	"time"

	"encore.app/content/queries"
)

// Post is an entry in the blog.
type Post struct {
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
	Post Post `json:"post"`
}

// CreatePost adds a new post to the database.
//
//encore:api public method=POST path=/content/posts
func CreatePost(ctx context.Context, params *CreatePostParams) (*CreatePostResponse, error) {
	p, err := queries.New(contentDB.Stdlib()).CreatePost(ctx, queries.CreatePostParams{
		Slug:    params.Slug,
		Title:   params.Title,
		Summary: params.Summary,
		Body:    params.Body,
	})
	if err != nil {
		return nil, err
	}

	return &CreatePostResponse{
		Post: Post{
			Slug:      p.Slug,
			Title:     p.Title,
			Summary:   p.Summary,
			Body:      p.Body,
			CreatedAt: p.CreatedAt,
		},
	}, nil
}

// GetPostResponse is the response data for the GetPosts endpoint.
type GetPostResponse struct {
	Post Post
}

// GetPost returns all posts.
//
//encore:api public method=GET path=/content/posts/:slug
func GetPost(ctx context.Context, slug string) (*GetPostResponse, error) {
	post, err := queries.New(contentDB.Stdlib()).GetPost(ctx, slug)
	if err != nil {
		return nil, err
	}

	return &GetPostResponse{
		Post: Post{
			Slug:      post.Slug,
			Title:     post.Title,
			Summary:   post.Summary,
			Body:      post.Body,
			CreatedAt: post.CreatedAt,
		},
	}, nil
}

// GetPostsResponse is the response data for the GetPosts endpoint.
type GetPostsResponse struct {
	// Collection of posts.
	Posts []Post
}

// GetPosts returns all posts.
//
//encore:api public method=GET path=/content/posts
func GetPosts(ctx context.Context) (*GetPostsResponse, error) {
	posts, err := queries.New(contentDB.Stdlib()).GetPosts(ctx)
	if err != nil {
		return nil, err
	}

	res := &GetPostsResponse{
		Posts: make([]Post, len(posts)),
	}

	for i, p := range posts {
		res.Posts[i] = Post{
			Slug:      p.Slug,
			Title:     p.Title,
			Summary:   p.Summary,
			Body:      p.Body,
			CreatedAt: p.CreatedAt,
		}
	}

	return res, nil
}
