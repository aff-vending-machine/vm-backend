package migration

import (
	"time"
	"vm-backend/internal/core/domain/catalog"
	"vm-backend/internal/core/domain/machine"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	SKU       string    `json:"sku" gorm:"uniqueIndex"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	ImageURL  string    `json:"image_url"`
	Price     float64   `json:"price"`
}

func (e Product) TableName() string {
	return "products"
}

func MigrateProduct(db *gorm.DB) {
	db.AutoMigrate(&Product{})

	var products []Product
	db.Find(&products)

	for _, product := range products {
		var catalogGroup catalog.Group

		if product.Type == "" {
			product.Type = "Uncategorized"
		}

		// Check if the role already exists in the new table
		if isNotFound(db.Where("name = ?", product.Type), &catalogGroup) {
			catalogGroup = catalog.Group{
				Name:        product.Type,
				Description: "",
				IsEnable:    true,
			}
			db.Create(&catalogGroup)
			log.Info().Str("name", catalogGroup.Name).Msg("migrated group")
		}

		var catalogProduct catalog.Product

		// Check if the user already exists in the new table
		if isNotFound(db.Where("sku = ?", product.SKU), &catalogProduct) {
			catalogProduct = catalog.Product{
				ID:           product.ID,
				GroupID:      catalogGroup.ID,
				CreatedAt:    product.CreatedAt,
				UpdatedAt:    product.UpdatedAt,
				SKU:          product.SKU,
				Name:         product.Name,
				Description:  "",
				ImageURL:     product.ImageURL,
				Barcode:      "",
				ProductPrice: product.Price,
				SalePrice:    product.Price,
				IsEnable:     true,
			}
			db.Create(&catalogProduct)
			log.Info().Str("sku", catalogProduct.SKU).Msg("migrated product")
			tx := db.Model(&machine.Slot{}).Where("product_id = ?", product.ID).Update("catalog_product_id", catalogProduct.ID)
			log.Info().Uint("product_id", product.ID).Int64("affected", tx.RowsAffected).Msg("migrated slots")
		}
	}
}
