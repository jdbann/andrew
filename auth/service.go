package auth

// Service holds configuration for the auth service. Used here to allow the api token to be set in test environments.
//
//encore:service
type Service struct {
	StableAPIToken string
}

func initService() (*Service, error) {
	return &Service{
		StableAPIToken: secrets.StableAPIToken,
	}, nil
}

var secrets struct {
	StableAPIToken string
}
