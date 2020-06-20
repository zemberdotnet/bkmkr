package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	jsonMap := app.getAllDocs("users")

	td := &templateData{
		People: jsonMap,
	}
	app.render(w, r, "home.page.tmpl", td)
	fmt.Println(r.Header)
}

func (app *application) clear(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	iter := newIter(ctx, app.db, "users")
	docs, err := iter.GetAll()
	// Work towards better error handling
	if err != nil {
		app.serverError(w, fmt.Errorf("Internal Server Error %v", http.StatusInternalServerError))
	}
	forms := make(Forms, len(docs))
	for i := range docs {
		data := docs[i].Data()
		fmt.Println(data)
		forms[i] = FormResp{docs[i].Ref.ID, data["born"].(int64), data["first"].(string), data["last"].(string)}
	}
	td := &templateData{
		Forms: &forms,
	}

	app.render(w, r, "clear.page.tmpl", td)
}

// room for improvement
func (app *application) del(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	fmt.Println(r.PostForm)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	ID := r.PostForm.Get("selection")
	ctx := context.Background()

	doc := app.db.Collection("users").Doc(ID)
	_, err = doc.Delete(ctx)
	if err != nil {
		app.serverError(w, err)
	}
	fmt.Fprint(w, "Document  Succesfully Deleted")
}

func (app *application) form(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, "create.page.tmpl", nil)
}

// Create relies on the input being non-null. We use the required attribute
// In the html document, but manual submission might bypass this
func (app *application) create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	first := r.PostForm.Get("first")
	last := r.PostForm.Get("last")
	year := r.PostForm.Get("year")
	yearconv, err := strconv.Atoi(year)
	if err != nil {
		app.serverError(w, err)
		return
	}
	m := make(map[string]interface{})
	m = map[string]interface{}{
		"first": first,
		"last":  last,
		"born":  yearconv,
	}

	ctx := context.Background()
	err = NewDoc(ctx, app.db, "users", m)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if r.Method == "POST" {
		http.Redirect(w, r, "/", 303)
	}
}
