package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

func (server *Server) Initialize(appConfig AppConfig) {
	fmt.Println("Welcome To " + appConfig.AppName)
	server.Router = mux.NewRouter()
	server.initializeRouter()
}

// untuk run servernya
func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

//validasi apakah ada file .env
func getEnv(key, fallback string) string  {
	if value, ok := os.LookupEnv(key)
}


func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error Loading .env.example file")
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppEnv = os.Getenv("APP_ENV")
	appConfig.AppPort = os.Getenv("APP_PORT")

	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
