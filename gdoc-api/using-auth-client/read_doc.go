package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

func main() {
	// Create a new HTTP client with the necessary OAuth 2.0 credentials.
	ctx := context.Background()
	client, err := getClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create HTTP client: %v", err)
	}

	// Create a new Google Docs service client.
	srv, err := docs.NewService(ctx, option.WithHTTPClient(client))
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
	log.Println(doc.Body)
}
func getClient(ctx context.Context) (*http.Client, error) {
	// Replace "credentials.json" with the path to your OAuth 2.0 credentials JSON file.
	//file, err := os.Open("/Users/reeta/secrets/client_secret_35544671621-h27rr6s79rfgst7vcsvgr2kkbht9geku.apps.googleusercontent.com.json")
	data, err := os.ReadFile("/Users/reeta/secrets/creds.json")
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	cred, err := google.CredentialsFromJSON(ctx, data, docs.DocumentsScope)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve OAuth 2.0 credentials: %v", err)
	}

	// Create an HTTP client with the necessary OAuth 2.0 credentials.
	client := oauth2.NewClient(ctx, cred.TokenSource)
	if client == nil {
		return nil, fmt.Errorf("clien is nil")
	}
	return client, nil
}
