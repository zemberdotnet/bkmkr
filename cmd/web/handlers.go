package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

var Forms = &FormArr{}

type HomePage struct {
	Home string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	ctx := context.Background()
	iter := app.getDocs(app.db, ctx, "users")

	jsonString, _ := json.MarshalIndent(Read(app.db, ctx, iter), "", "\t")
	s := string(jsonString)
	Home := HomePage{
		Home: s,
	}
	/*
		files := []string{
			"./ui/html/home.page.tmpl",
			"./ui/html/footer.partial.tmpl",
			"./ui/html/base.layout.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, err)
			return
		}
	*/
	err := app.templateCache["home.page.tmpl"].Execute(w, Home)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) clear(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	iter := app.getDocs(app.db, ctx, "users")
	options := Read(app.db, ctx, iter)
	iter = app.getDocs(app.db, ctx, "users")
	docs := GetDocs(app.db, ctx, iter)

	formResp := make([]FormResp, 0)

	for i, e := range options {

		formResp = append(formResp, FormResp{i, docs[i].Ref, e["born"].(int64), e["first"].(string), e["last"].(string)})

	}
	//global might not be the way here
	Forms = &FormArr{
		Forms: formResp,
	}

	// Okay so I may be destroyging memory here by appending testMarsh

	// I dont know what the proper type to append
	// a new array,a pointer, a reference???
	//testArr = append(testArr, testMarsh{

	files := []string{
		"./ui/html/del.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, formResp)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) del(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	delDoc := r.PostForm.Get("selection")
	//  Simplify this down, the format is more complex than needed
	docIndex, err := strconv.Atoi(delDoc)
	ctx := context.Background()
	_, err = Forms.Forms[docIndex].DocRef.Delete(ctx)
	if err != nil {
		app.serverError(w, err)
	}
	fmt.Fprint(w, "Document Deleted")
}

func (app *application) form(w http.ResponseWriter, r *http.Request) {

	// The order of the files is suprisingly important
	files := []string{
		"./ui/html/footer.partial.tmpl",
		"./ui/html/create.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, "No data needed")
	if err != nil {
		app.serverError(w, err)
	}
}

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
	err = NewEntry("users", m, app.db, ctx)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprint(os.Stdout, first, last, year)
	app.home(w, r)
}
