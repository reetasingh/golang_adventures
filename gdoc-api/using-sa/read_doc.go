package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

func main() {
	// Create a new HTTP client with the necessary OAuth 2.0 credentials.
	ctx := context.Background()

	// Create a new Google Docs service client.
	srv, err := docs.NewService(ctx, option.WithServiceAccountFile("/Users/reeta/secrets/serviceaccount.json"))
	if err != nil {
		log.Fatalf("Failed to create Google Docs service client: %v", err)
	}

	// Use the srv object to make API calls to Google Docs.
	// For example, you can list the documents in the user's Google Drive.
	documentid, err := os.ReadFile("../document-id.txt")
	doc, err := srv.Documents.Get(string(documentid)).Do()
	if err != nil {
		log.Fatalf("Failed to fetch document: %v", err)
	}
	log.Println(doc.Title)
}
