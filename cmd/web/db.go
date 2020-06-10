package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"reflect"
)

func connectInit(ctx context.Context) (*firestore.Client, error) {
	sa := option.WithCredentialsFile("/home/user1/Downloads/MyFirstProject.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		fmt.Println(1)
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client, err
}

func connectClose(client *firestore.Client) {
	client.Close()
}

func NewEntry(coll string, m map[string]interface{}, client *firestore.Client, ctx context.Context) error {
	_, _, err := client.Collection(coll).Add(ctx, m)
	if err != nil {
		log.Fatalf("Failed adding: %v", err)
	}
	return err
}

// Iter is a bad dependency we should shoot for a more general type of query
// Okay so low key this feels like i just coppied whatever iter is
func GetDocs(c *firestore.Client, ctx context.Context, iter *firestore.DocumentIterator) []*firestore.DocumentSnapshot {
	// a bad way to do this but whatever / Dynamism we will be appending to an empty array
	docs := make([]*firestore.DocumentSnapshot, 0)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			log.Fatal("There has been an error: %v", err)
		}
		fmt.Println("DOCS------------------------")
		fmt.Println(doc)
		docs = append(docs, doc)
	}
	return docs
}

func Read(c *firestore.Client, ctx context.Context, iter *firestore.DocumentIterator) []map[string]interface{} {
	// set len/capacity to 1000, but need to be more dynamic
	testArr := make([]map[string]interface{}, 0, 1000)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			log.Fatal("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
		// FIX THIS NAME BETTER

		testArr = append(testArr, doc.Data())

	}
	return testArr
}

func (app *application) delAll(c *firestore.Client, ctx context.Context, iter *firestore.DocumentIterator) {
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			log.Fatal("Failed to iterate: %v", err)
		}
		fmt.Println(reflect.TypeOf(doc.Ref).String())
		_, err = doc.Ref.Delete(ctx)
		if err != nil {
			log.Fatal("Failed to delete: %v", err)
		}
	}
}

//Doesn't work we need to understand the document concept better
/*
func (app *application) testFunc (client *firestore.Client, ctx context.Context) {
	docs, err := client.GetAll(ctx, []*firestore.DocumentRef {
		client.Doc("users"),
	})
	if err != nil {
		app.errorLog.Output(2, "not good")
	}
	fmt.Println(docs)
}
*/
// How would this work if there was a sub collection of documents
// Gaining understanding on how this works. client.... gives us a document iterator struct
// which has fields iter - document Iterator and err - error
// docIterator is a interface that defines method next() and stop()
// its an interface so you can have straight queries and one for query snapshots

func (app *application) getDocs(client *firestore.Client, ctx context.Context, coll string) *firestore.DocumentIterator {
	return client.Collection(coll).Documents(ctx)
}

func (app *application) testColl(client *firestore.Client, ctx context.Context) {
	coll1 := client.Collection("users")
	fmt.Println("START TESTING")
	fmt.Println(reflect.TypeOf(coll1).String())
	fmt.Println(coll1)
}
