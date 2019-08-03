package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	projectID := os.Getenv("PROJECT_ID")
	// Get a Firestore client.
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close client when done.
	defer client.Close()

	fmt.Println("Hello World")
}
