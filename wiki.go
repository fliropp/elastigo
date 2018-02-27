package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title  string
	Body   string
	Author string
	Id     string
}

/*func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}*/

func viewHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/view/"):]
	client := getClient()
	entry, err := getWikiEntry(client, id)
	if err != nil {
		http.Redirect(w, r, "/edit/"+id, http.StatusFound)
		return
	}
	p := Page{Title: entry.Title, Body: entry.Body, Author: entry.User, Id: id}
	t, _ := template.ParseFiles("html/view.html")
	t.Execute(w, p)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/edit/"):]
	client := getClient()
	entry, err := getWikiEntry(client, id)
	p := &Page{Title: "add title...", Body: "add body...", Author: "add user...", Id: id}
	if err == nil {
		p = &Page{Title: entry.Title, Body: entry.Body, Author: entry.User, Id: id}
	}
	t, _ := template.ParseFiles("html/edit.html")
	t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/save/"):]
	title := r.FormValue("title")
	body := r.FormValue("body")
	author := r.FormValue("author")

	//p := &Page{Title: title, Body: body, Author: author, Id: id}
	addWikiEntry(getClient(), WikiEntry{User: author, Body: body, Title: title}, id)
	http.Redirect(w, r, "/view/"+id, http.StatusFound)
}
