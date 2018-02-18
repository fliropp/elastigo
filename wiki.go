package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title  string
	Body   []byte
	Author string
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
		http.Redirect(w, r, "/edit/"+entry.Title, http.StatusFound)
		return
	}
	body, _ := ioutil.ReadFile(entry.Body)
	p := Page{Title: entry.Title, Body: body, Author: entry.User}
	t, _ := template.ParseFiles("html/view.html")
	t.Execute(w, p)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/edit/"):]
	client := getClient()
	entry, err := getWikiEntry(client, id)

	p := &Page{Title: "add title...", Body: []byte("add body..."), Author: "add user..."}
	if err == nil {
		body, _ := ioutil.ReadFile(entry.Body)
		p = &Page{Title: entry.Title, Body: body, Author: entry.User}
	}
	t, _ := template.ParseFiles("html/edit.html")
	t.Execute(w, p)
}

/*func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}*/
