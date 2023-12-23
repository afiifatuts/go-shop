package fakers

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/afiifatuts/go-shop/app/models"
	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ProductFaker(db *gorm.DB) *models.Product {

	//karena di table product ada relasi ke table user
	//maka kita harus membuat fake productnya terlebih dahulu
	user := UserFaker(db)
	err := db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	name := faker.Name()
	return &models.Product{
		ID:               uuid.New().String(),
		UserID:           user.ID,
		Sku:              slug.Make(name),
		Name:             name,
		Slug:             slug.Make(name),
		Price:            decimal.NewFromFloat(fakePrice()),
		Stock:            rand.Intn(100),
		Weight:           decimal.NewFromFloat(rand.Float64()),
		ShortDescription: faker.Paragraph(),
		Description:      faker.Paragraph(),
		Status:           0,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}
}

//membuat method untuk generate pricenta

func fakePrice() float64 {
	return precision(rand.Float64()*math.Pow10(rand.Intn(8)), rand.Intn(2)+1)
}

func precision(val float64, pre int) float64 {
	div := math.Pow10(pre)
	return float64(int64(val*div)) / div
}
