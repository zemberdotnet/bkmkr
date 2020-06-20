package main

import (
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	db            *firestore.Client
	templateCache map[string]*template.Template
}

// Working towards renaming client in application struct or just creating a
// new pattern
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO:\t ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERR:\t ", log.Ldate|log.Ltime|log.Lshortfile)

	ctx := context.Background()
	creds := "/home/user1/Downloads/MyFirstProject.json"
	c, err := firestore.NewClient(ctx, firestore.DetectProjectID, option.WithCredentialsFile(creds))
	if err != nil {
		errorLog.Fatalf("New Client: %v", err)
	}
	ts, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		db:            c,
		templateCache: ts,
	}

	infoLog.Printf("Starting server on %s", *addr)
	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}
	// Testing
	/**

	type TestStruct struct {
		first string
		last  string
		year  int
	}

	group  := TestStruct{
		first
	**/
	//NewEntry("testcoll", client)
	//	m := make(map[string]interface{})
	//	m = map[string]interface{}{
	//		"first":"Jay",
	//		"last":"Minor",
	//		"born":2100,
	//	}
	//	err = NewEntry("users", m, client, ctx)

	// This was working pretty well for json stuff
	/*
		x := Read(client, ctx)
		jsonString, _ := json.Marshal(x)
		fmt.Println(string(jsonString))
	*/

	//	Tests for a users document which doesnt exists...users is our collection
	//	app.testFunc(client, ctx)
	//	app.testColl(client, ctx)
	//	iter := app.getDocs(client, ctx, "users")
	//	x := Read(client, ctx, iter)
	//	jsonString, _ := json.Marshal(x)
	//	fmt.Println(string(jsonString))

	/*
			This is Set.
		testData := make(map[string]interface{})
		testData["Born"] = 1998
		coll := client.Collection("users")
		docref := coll.Doc("0001")
		_, err = docref.Set(ctx, testData)
	*/
	err = srv.ListenAndServe()

	log.Fatal(err)

}
