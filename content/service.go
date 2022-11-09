package content

import (
	"encore.dev/config"
	"encore.dev/storage/sqldb"
)

var contentDB = sqldb.Named("content")

// ServiceCfg holds configuration for the content service.
type ServiceCfg struct {
	// If true, prime DB with useful test content on launch.
	GenerateSeeds config.Bool
}

var cfg *ServiceCfg = config.Load[*ServiceCfg]()

func init() {
	if cfg.GenerateSeeds() {
		generateSeeds()
	}
}
