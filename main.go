package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// URL struct (Model)
type URL struct {
	gorm.Model
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

func main() {
	enverr := godotenv.Load()
if enverr!=nil{
	fmt.Println("Error loading environment variables")
}
    postgres_uri:=os.Getenv("DATABASE_URI")
	// Define PostgreSQL connection string
	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(postgres_uri), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} else {
		fmt.Println("Database connected")
	}

	// AutoMigrate creates the table if it doesn't exist
	db.AutoMigrate(&URL{})

	// Set up HTTP handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there")  // Using Fprintf instead of Println
	})
	
	// Start the server (this will block until the server errors)
	fmt.Println("server is running on http://localhost:8080")
	serveErr := http.ListenAndServe(":8080", nil)
	if serveErr!=nil {
		log.Fatal(err)
	}
}