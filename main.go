package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// It's saves the template of our site in the cache to improve
// performance and save memory. Now we don't need to do this every
// time the function 'getPage' is called.
var tmpl = template.Must(template.ParseFiles("assets/index.html"))

type Post struct {
	Text string
	Time string
}

// Main struct.
type Page struct {
	Posts []Post
}

// All our posts saves in this slice.
var allPosts []Post

func getPage(w http.ResponseWriter, req *http.Request) {
	tmpl.Execute(w, Page{
		Posts: allPosts,
	})
}

func createPost(w http.ResponseWriter, req *http.Request) {
	// Read text from textarea in html template.
	text := req.FormValue("newpost")

	if text != "" {
		newPost := Post{
			Text: text,
			Time: time.Now().Format("2 Jan 15:04"),
		}

		// Insert a new post in start of our slice.
		var p []Post
		p = append(p, newPost)
		allPosts = append(allPosts[:0], append(p, allPosts[0:]...)...)
	}

	http.Redirect(w, req, "/", http.StatusFound)
}

func main() {
	r := mux.NewRouter()

	// Connect css-file in directory assets/style.
	fs := http.FileServer(http.Dir("assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	r.HandleFunc("/", getPage).Methods("GET")
	r.HandleFunc("/", createPost).Methods("POST")

	http.ListenAndServe(":8080", r)
}
