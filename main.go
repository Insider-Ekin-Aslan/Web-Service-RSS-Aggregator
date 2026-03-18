package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Insider-Ekin-Aslan/Web-Application-RSS-Aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not found in the environment.")
	}

	println("Running at PORT:", port)

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not found in the environment.")
	}

	println("Database URL found. DATABASE_URL:", databaseURL)

	connection, connectionError := sql.Open("postgres", databaseURL)

	if connectionError != nil {
		log.Fatal("Can't connect to database. ERROR:", connectionError)
	}

	config := apiConfig{Database: database.New(connection)}

	// config.Database.CreateUser(context.Background(), database.CreateUserParams{
	// 	ID:        uuid.New(),
	// 	CreatedAt: time.Now().UTC(),
	// 	UpdatedAt: time.Now().UTC(),
	// 	Name:      "test",
	// })

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routerV1 := chi.NewRouter()

	routerV1.Get("/healthz", handlerReadiness)
	routerV1.Get("/error", handlerError)
	routerV1.Post("/users", config.handlerCreateUser)

	router.Mount("/v1", routerV1)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	serverError := server.ListenAndServe()

	if serverError != nil {
		log.Fatal(serverError)
	}
}

type apiConfig struct {
	Database *database.Queries
}
