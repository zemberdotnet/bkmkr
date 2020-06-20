package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func connectClose(client *firestore.Client) {
	client.Close()
}

func NewDoc(ctx context.Context, client *firestore.Client, coll string, m map[string]interface{}) error {
	_, _, err := client.Collection(coll).Add(ctx, m)
	if err != nil {
		log.Fatalf("Failed adding: %v", err)
	}
	return err
}

// How would this work if there was a sub collection of documents
// Gaining understanding on how this works. client.... gives us a document iterator struct
// which has fields iter - document Iterator and err - error
// docIterator is a interface that defines method next() and stop()
// its an interface so you can have straight queries and one for query snapshots

// I wonder if writing this as a method would be better. Then maybe it could be
// adaptable across different applications??
func newIter(ctx context.Context, client *firestore.Client, coll string) *firestore.DocumentIterator {
	return client.Collection(coll).Documents(ctx)
}

func (app *application) getAllDocs(collection string) *[]map[string]interface{} {
	ctx := context.Background()
	iter := newIter(ctx, app.db, collection)
	docs, err := iter.GetAll()
	jsonResp := make([]map[string]interface{}, len(docs))
	if err != nil {
		app.errorLog.Fatalf("Failed to read docs: %v", err)
	}
	for d := range docs {
		jsonResp[d] = docs[d].Data()
	}
	return &jsonResp
}
