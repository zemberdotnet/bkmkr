package main

import (
	"fmt"
	"net/http"	
	"context"
	//"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	ctx := context.Background()
	fmt.Fprint(w, Read(app.db, ctx))
}
