package repository

import (
	"github.com/afiifatuts/go-shop/app/models"
	"gorm.io/gorm"
)

type Product struct {
	*models.Product
}

func (p *Product) GetProducts(db *gorm.DB) (*[]models.Product, error) {
	var err error
	var products []models.Product
	err = db.Debug().Model(&models.Product{}).Limit(20).Find(&products).Error

	if err != nil {
		return nil, err
	}
	return &products, nil
}
