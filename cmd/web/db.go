package main

import (
	"context"
	"google.golang.org/api/iterator"
	"os"
	"fmt"
	"log"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"reflect"
	"cloud.google.com/go/firestore"
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

func NewEntry(coll string, client *firestore.Client) {
	collref := client.Collection(coll)
	fmt.Println(collref)
}

	/*
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first":"Ada",
		"last":"Lovelace",
		"born": 1815,
	})
	
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	*/

func Read(c *firestore.Client, ctx context.Context) string {
	iter := c.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			//iter.Stop()
			// Don't think the above is necessary. When it reaches the end iter.Stop is already been called, maybe maybe not we will see
			break

		}
		if err != nil {
			log.Fatal("Failed to iterate: %v", err)
		}
		//return doc.Data()
		x := doc.Data()
		for k, v := range x {
			fmt.Println("k:", k, "v:", v)	
		}
		fmt.Println(x["first"])
		fmt.Fprintf(os.Stdout,"Data: %v\n", doc.Data())

		fmt.Fprintf(os.Stdout, "Type %v\n", reflect.TypeOf(doc.Data()).String())
			
	}
	return ""
}

