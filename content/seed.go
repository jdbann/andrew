package content

import (
	"context"
	"time"

	"encore.app/content/queries"
	"encore.dev"
	"encore.dev/rlog"
)

func generateSeeds() {
	// Double check that we're not seeding production.
	if encore.Meta().Environment.Type == encore.EnvProduction {
		rlog.Error("attempted to seed database in production environment")
		panic("attempted to seed database in production environment")
	}

	rlog.Info("seeding database")

	posts := []queries.SeedPostParams{{
		Slug:      "designing-as-a-developer",
		Title:     "Designing as a developer",
		Summary:   "Design always felt like a vague and unknowable discipline to me. That it sat in the world of ambiguous creation and required a certain innate confidence that 'what I have designed is what is required'. That's not how I work as a developer, in a land of unit tests and exception tracking.",
		Body:      "Design always felt like a vague and unknowable discipline to me. That it sat in the world of ambiguous creation and required a certain innate confidence that 'what I have designed is what is required'. That's not how I work as a developer, in a land of unit tests and exception tracking.\n\nWhat I'm missing is some way of verifying that my work is doing what it is meant to do. I thought that in the world of design this was some innate aesthetic judgement, and that it took confidence and experience to 'just know' that the design was correct.",
		CreatedAt: time.Date(2022, time.September, 5, 18, 45, 0, 0, time.UTC),
	}, {
		Slug:      "learning-to-love-tailwind",
		Title:     "Learning to love Tailwind",
		Summary:   "I wasn't always a fan, but over the last year I've become increasingly won over by a jumble of classes. But the sprinkling of extra magic you can get from really tearing into the config file makes for a very elegant interface for building styles without leaving your markup.",
		Body:      "I wasn't always a fan, but over the last year I've become increasingly won over by a jumble of classes. But the sprinkling of extra magic you can get from really tearing into the config file makes for a very elegant interface for building styles without leaving your markup.",
		CreatedAt: time.Date(2022, time.September, 4, 16, 55, 0, 0, time.UTC),
	}}

	conn := queries.New(contentDB.Stdlib())

	for _, post := range posts {
		if err := conn.SeedPost(context.Background(), post); err != nil {
			rlog.Error("unable to seed", "err", err)
			panic(err)
		}
	}
}
