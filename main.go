package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/hiteshVasisht99/excel-import/handlers"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbClient := EstablishDBConnection()
	redisClient := EstablishRedisConn()

	handlers.InitializeService(dbClient, redisClient)

	r := mux.NewRouter()

	r.HandleFunc("/employee/create", handlers.UploadFile).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", handlers.GetEmployeeDataByID).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", handlers.UpdateEmployeeDataByID).Methods(http.MethodPut)

	log.Println("Starting the service")
	http.ListenAndServe(":8108", r)

}

func EstablishDBConnection() *sql.DB {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dbConnString)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	log.Println("Connection Established Successfully")
	return db
}

func EstablishRedisConn() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	client := redis.NewClient(&redis.Options{
		Addr: redisHost,
		DB:   0,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Redis connection established successfully :", pong)
	return client
}
