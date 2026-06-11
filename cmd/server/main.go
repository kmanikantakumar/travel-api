package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"travel-api/internal/cab"
	"travel-api/internal/flights"
	"travel-api/internal/hotels"
	"travel-api/internal/intent"
	"travel-api/internal/rail"
	"travel-api/internal/trips"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Println("Warning: .env file not found or error loading it:", err)
	}
	http.HandleFunc("/api/search-flights", flights.SearchHandler)
	http.HandleFunc("/api/intent", intent.Handler)
	http.HandleFunc("/api/search-rails", rail.SearchRailsHandler)
	http.HandleFunc("/api/cab/book", cab.BookCabHandler)
	http.HandleFunc("/api/hotels/search", hotels.SearchHotelsHandler)

	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println("DB URL:", dbURL)

	if dbURL == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	defer conn.Close(context.Background())

	fmt.Println("Connected!")

	// Create a sample trip
	tripReq := trips.CreateTripRequest{}

	trip, err := trips.Create(conn, tripReq)
	if err != nil {
		fmt.Println("Error creating trip:", err)
	} else {
		fmt.Printf("Trip created: %+v\n", trip)
	}

	http.HandleFunc("/api/trips", trips.CreateTripHandler(conn))
	http.HandleFunc("/api/trips/{id}", trips.GetTripHandler(conn))

	fmt.Println("Connected!")
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello world")
// }

// func main2() {
// 	godotenv.Load()

// 	dbURL := os.Getenv("DATABASE_URL")

// 	fmt.Println("DB URL:", dbURL)

// 	conn, err := pgx.Connect(context.Background(), dbURL)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer conn.Close(context.Background())

// 	fmt.Println("Connected!")

// 	from := mail.NewEmail("Example User", "keralaflashman@gmail.com")
// 	subject := "Sending with SendGrid is Fun"
// 	to := mail.NewEmail("Example User", "kumar2209b@gmail.com")
// 	plainTextContent := "and easy to do anywhere, even with Go"
// 	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
// 	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
// 	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
// 	// client.Request, _ = sendgrid.SetDataResidency(client.Request, "eu")
// 	// uncomment the above line if you are sending mail using a regional EU subuser
// 	response, err := client.Send(message)
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		fmt.Println(response.StatusCode)
// 		fmt.Println(response.Body)
// 		fmt.Println(response.Headers)
// 	}

// }
