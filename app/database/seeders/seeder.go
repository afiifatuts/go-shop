package seeders

import (
	"github.com/afiifatuts/go-shop/app/database/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

// memesukkan data faker ke seeder
func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.UserFaker(db)},
		{Seeder: fakers.ProductFaker(db)},
	}
}

// untuk proses insert ke databasenya
// jadi nantinya dia looping faker yang telah dibuat
// kemudian dimasukkan ke database
func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}
	return nil
}
