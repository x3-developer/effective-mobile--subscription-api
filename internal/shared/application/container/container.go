package container

import (
	"subscriptions/config"
	"subscriptions/internal/shared/persistence"
)

type Container struct {
	PostgresDB *persistence.Postgres
}

func NewContainer(cfg *config.Config) *Container {
	pdb := persistence.NewPostgres(cfg)

	return &Container{
		PostgresDB: pdb,
	}
}
