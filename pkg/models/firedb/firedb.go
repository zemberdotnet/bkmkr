package firedb 

import (
	"context"
	"fmt"
	"log"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)
 
// Context needs to go to a struct
// Need to use the Struct we built in models in WriteDB?

//Probably something along the lines of having a struct for User
//Then having a sub struct in user with the records they have?
// User -> Records -> Individual Records
func InitApp(keyFile string) (*firestore.Client, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile(keyFile)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	
	return client, err 
}

func (a *firestore.Client) WriteDB(coll string,  ) {
	_, _, err = client.Collection(coll).Add(ctx
