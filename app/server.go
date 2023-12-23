package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/afiifatuts/go-shop/app/database/seeders"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

// membuat struct
type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

// parameter initialize adalah value yang ada di function Run()
func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome To " + appConfig.AppName)
	//make connection to database

	server.initializeDb(dbConfig)
	server.initializeRouter()
	seeders.DBSeed(server.DB)
}

// supaya lebih rapi membuat method intialize db secara terpisah
func (server *Server) initializeDb(dbConfig DBConfig) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed on connecting to the database server")
	}

	//migration db dan menangkap data dari registry.go
	//lakukan looping dulu untuk setiap models
	for _, model := range RegisterModels() {
		err = server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migration successfully")

}

// untuk run servernya
func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

// validasi apakah ada file .env
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		// kalau ketemu key nya dia akan return valuenya
		return value
	}

	// kalau tidak ditemukan keynya maka return fallbacknya
	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error Loading .env.example file")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoToko")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	//asign value dari struct yang telah dibuat
	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "postgres")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "blimbeng38")
	dbConfig.DBName = getEnv("DB_NAME", "goshop_2023")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
