package repositories

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/adapters/db"
	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/adapters/db/config"
	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/core/models"
	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/core/service"
	"gorm.io/gorm"
)

var Tenants = []models.Tenant{
	{Name: "Beyaz Futbol", Slug: "beyaz_futbol", Port: 5432},
	{Name: "Hell Kitchen", Slug: "hell_kitchen", Port: 5433},
	{Name: "Mentalist", Slug: "mentalist", Port: 5434},
}

type DBMultiplexer struct {
	clients map[string]*gorm.DB
}

func NewDBMultiplexer() (*DBMultiplexer, error) {

	clients := make(map[string]*gorm.DB)
	for _, tenant := range Tenants {

		config := config.NewDatabaseConfig(tenant.Slug, tenant.Port)
		client, err := db.ConnectPostgres(config)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to postgres for tenant %s: %w", tenant.Name, err)
		}

		clients[tenant.Slug] = client
	}

	return &DBMultiplexer{
		clients: clients,
	}, nil
}

func (mux *DBMultiplexer) GetClient(ctx context.Context) (*gorm.DB, error) {
	dbName, err := service.GetDb(ctx)
	if err != nil {
		return nil, fmt.Errorf("tenant not found")
	}

	client, ok := mux.clients[dbName]
	if !ok {
		return nil, fmt.Errorf("tenant not found")
	}
	return client, nil
}
