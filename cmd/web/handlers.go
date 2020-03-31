package main

import (
	"fmt"
	"net/http"	
	//"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	fmt.Println("Test")
}
