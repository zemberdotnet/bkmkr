package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	db            *firestore.Client
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO:\t ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERR:\t ", log.Ldate|log.Ltime|log.Lshortfile)

	ctx := context.Background()
	client, err := connectInit(ctx)
	defer connectClose(client)

	ts, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		db:            client,
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

	err = srv.ListenAndServe()

	log.Fatal(err)

}
