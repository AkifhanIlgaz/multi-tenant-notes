package repositories

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/service"
	"gorm.io/gorm"
)

type SchemaWrapper struct {
	db *gorm.DB
}

func NewSchemaWrapper(db *gorm.DB) *SchemaWrapper {
	return &SchemaWrapper{db: db}
}

func (sw *SchemaWrapper) CreateSchema(schemaName string) error {
	return sw.db.Exec("CREATE SCHEMA IF NOT EXISTS ?", schemaName).Error
}

func (sw *SchemaWrapper) ExecuteWithSchema(ctx context.Context, fn func(*gorm.DB) error) error {
	schema, err := service.GetSchema(ctx)
	if err != nil {
		return err
	}

	return sw.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(fmt.Sprintf("SET LOCAL search_path TO %s", schema)).Error; err != nil {
			return err
		}

		return fn(tx.WithContext(ctx))
	})
}
