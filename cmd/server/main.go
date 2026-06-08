package main

import (
	"fmt"
	"net/http"
	"travel-api/internal/flights"
	"travel-api/internal/intent"
	"travel-api/internal/rail"
)

func main() {
	http.HandleFunc("/api/search-flights", flights.SearchHandler)
	http.HandleFunc("/api/intent", intent.Handler)
	http.HandleFunc("/api/search-rails", rail.SearchRailsHandler)

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
