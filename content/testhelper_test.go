package content_test

import (
	"context"
	"testing"

	"encore.dev/storage/sqldb"
)

var contentDB = sqldb.Named("content")

func setupDB(t *testing.T) {
	if err := truncateDB(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := truncateDB(); err != nil {
			t.Fatal(err)
		}
	})
}

func truncateDB() error {
	_, err := contentDB.Exec(context.Background(), "TRUNCATE posts")
	return err
}
