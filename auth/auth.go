package auth

import (
	"context"

	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
)

// AuthHandler is the method used by encore to authenticate a request based on a provided bearer token. The actual implementation is extracted here as the encore runtime does not allow the auth handler to be called from a separate package.
//
//encore:authhandler
func (s *Service) AuthHandler(ctx context.Context, token string) (auth.UID, error) {
	return s.TokenAuthHandler(ctx, token)
}

// TokenAuthHandler verifies that the correct bearer token has been provided.
func (s *Service) TokenAuthHandler(ctx context.Context, token string) (auth.UID, error) {
	if token != s.StableAPIToken {
		return "", &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "invalid token",
		}
	}

	return "stable", nil
}
