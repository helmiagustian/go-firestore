package firestore

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

var (
	// Articles collection name
	Articles = "articles"
)

// Add test add firestore
func Add(ctx context.Context, client *firestore.Client, col string, doc *map[string]interface{}) {
	_, _, err := client.Collection(col).Add(ctx, *doc)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}

// Read test read firestore
func Read(ctx context.Context, client *firestore.Client, col string) {
	iter := client.Collection(col).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}
