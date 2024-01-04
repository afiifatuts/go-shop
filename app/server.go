package app

import (
	"flag"
	"log"
	"os"

	"github.com/afiifatuts/go-shop/app/controllers"
	"github.com/joho/godotenv"
)

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
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

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

	// server.Initialize(appConfig, dbConfig)
	// server.Run(":" + appConfig.AppPort)

	//menangkap argumen setelah go run main {db:migrate}
	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.InitCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(": " + appConfig.AppPort)
		// log.Fatal(err)

	}
}
