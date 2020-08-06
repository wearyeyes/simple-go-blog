package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// It's saves the template of our site in the cache to improve
// performance and save memory.
var tmpl = template.Must(template.ParseFiles("assets/index.html"))

type Post struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Text string             `json:"text" bson:"text"`
	Time string             `json:"time" bson:"time"`
}

// Main struct.
type Page struct {
	Posts []Post
}

func main() {
	// Create router...
	r := mux.NewRouter()

	// Connect css-file in directory assets/style.
	fs := http.FileServer(http.Dir("assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	// Handlers.
	r.HandleFunc("/", getAllPosts).Methods("GET")
	r.HandleFunc("/", createNewPost).Methods("POST")

	log.Println("Server start working...")

	http.ListenAndServe(":8080", r)
}
