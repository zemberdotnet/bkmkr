package main

import (
	"context"
	"flag"
	"log"	
	"net/http"
	"os"
	"cloud.google.com/go/firestore"
)

type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
	db	 *firestore.Client
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	
	infoLog := log.New(os.Stdout, "INFO:\t ", log.Ldate | log.Ltime)
	errorLog := log.New(os.Stderr, "ERR:\t ", log.Ldate | log.Ltime | log.Lshortfile)

	ctx := context.Background()
	client, err := connectInit(ctx)
	defer connectClose(client)
	
	app := application{
		infoLog: infoLog,
		errorLog: errorLog,
		db: 	client,
	}

	infoLog.Printf("Starting server on %s", *addr)
	srv := &http.Server{
		Addr: *addr,
		Handler: app.routes(),
		ErrorLog: errorLog,
		}	

	//NewEntry("testcoll", client)
	Read(client, ctx)
	err = srv.ListenAndServe()
	log.Fatal(err)

}
