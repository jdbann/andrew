// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package queries

import (
	"time"
)

type Post struct {
	ID        int64
	Slug      string
	Title     string
	Summary   string
	Body      string
	CreatedAt time.Time
}