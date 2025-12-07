package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/KonstantinDuvakin/bd_app/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		// will exit the program and print the message
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("$PORT must be set")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("$dbUrl must be set")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	apiConf := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routerV1 := chi.NewRouter()
	routerV1.Get("/ready", handlerReadiness)
	routerV1.Get("/err", handlerError)
	routerV1.Post("/user", apiConf.createUser)
	routerV1.Get("/user", apiConf.getUser)

	router.Mount("/v1", routerV1)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Println("Starting server on port " + portString)
	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
