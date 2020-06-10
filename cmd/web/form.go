package main

import "cloud.google.com/go/firestore"

type FormArr struct {
	Forms []FormResp
}

type FormResp struct {
	ID     int
	DocRef *firestore.DocumentRef
	Born   int64
	First  string
	Last   string
}
