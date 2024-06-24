package main

import (
	"TODO-APP/pkg/handler"
	"TODO-APP/pkg/repository"
	"TODO-APP/pkg/service"
	"log"
	"os"

	"github.com/AntonZatsepilin/todo-app"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error init configs: %s", err.Error())
	}

	if err := godotenv.Load("/Users/anton/Desktop/GitHub/todo-app/.env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewReposytory(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
