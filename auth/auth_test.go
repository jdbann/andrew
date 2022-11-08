package auth_test

import (
	"context"
	"testing"

	"encore.app/auth"
	enc_auth "encore.dev/beta/auth"
	"encore.dev/beta/errs"
)

func TestTokenAuthHandler(t *testing.T) {
	s := auth.Service{
		StableAPIToken: "fake_api_token",
	}

	tests := []struct {
		name        string
		token       string
		wantUID     enc_auth.UID
		wantErrCode errs.ErrCode
	}{
		{
			name:        "success",
			token:       "fake_api_token",
			wantUID:     "stable",
			wantErrCode: errs.OK,
		},
		{
			name:        "failure",
			token:       "incorrect_api_token",
			wantUID:     "",
			wantErrCode: errs.Unauthenticated,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			uid, err := s.TokenAuthHandler(context.Background(), tt.token)

			if tt.wantErrCode != errs.Code(err) {
				t.Errorf("want error code %q; got %q", tt.wantErrCode.String(), errs.Code(err).String())
			}

			if tt.wantUID != uid {
				t.Errorf("want uid %q; got %q", tt.wantUID, uid)
			}
		})
	}
}
