package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/afiifatuts/go-shop/app/database/seeders"
	"github.com/afiifatuts/go-shop/app/models"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
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

	// server.initializeDb(dbConfig)-> tidak perlu karna kita harus tangkap argument dulu
	server.initializeDb(dbConfig)
	server.InitializeRouter()
	// seeders.DBSeed(server.DB)
}

// supaya lebih rapi membuat method intialize db secara terpisah
func (server *Server) initializeDb(dbConfig DBConfig) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed on connecting to the database server")
	}

}

func (server *Server) dbMigrate() {
	//migration db dan menangkap data dari registry.go
	//lakukan looping dulu untuk setiap models
	for _, model := range models.RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migration successfully")

}

// ini method untuk mendefinisikan command saat run go
func (server *Server) InitCommands(config AppConfig, db DBConfig) {
	// panggil db nya
	server.initializeDb(db)
	cmdApp := cli.NewApp()
	cmdApp.Commands = []*cli.Command{
		//devinisikan commandnya + methodnya
		{
			Name: "db:migrate",
			Action: func(ctx *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(ctx *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

// untuk run servernya
func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
