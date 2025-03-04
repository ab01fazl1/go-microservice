package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ab01fazl1/go-microservice/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool
	GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error)
	GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error)
	GetAllServices(ctx context.Context) ([]models.Service, error)
	GetAllVendors(ctx context.Context) ([]models.Vendor, error)
}

type Client struct {
	DB *gorm.DB
}


func (c Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}
func NewDatabaseClient() (DatabaseClient, error) {
	// hardcode, not good
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", "localhost",
		"postgres", "mysecretpassword", "postgres", 5432, "disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		DB: db,
	}
	return client, nil
}
