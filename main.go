package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")

	fmt.Println("DB URL:", dbURL)

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	fmt.Println("Connected!")

	from := mail.NewEmail("Example User", "keralaflashman@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "kumar2209b@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	// client.Request, _ = sendgrid.SetDataResidency(client.Request, "eu")
	// uncomment the above line if you are sending mail using a regional EU subuser
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

}
