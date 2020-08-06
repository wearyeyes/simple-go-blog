package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func getAllPosts(w http.ResponseWriter, req *http.Request) {
	var allPosts []Post

	collection := Connect()
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var onePost Post
		cursor.Decode(&onePost)

		// Add each value to the beginning of the slice.
		var helpSlice []Post
		helpSlice = append(helpSlice, onePost)

		allPosts = append(allPosts[:0], append(helpSlice, allPosts...)...)
	}

	tmpl.Execute(w, Page{
		Posts: allPosts,
	})
}

func createNewPost(w http.ResponseWriter, req *http.Request) {
	// Read text from textarea in html template.
	text := req.FormValue("newpost")

	if text != "" {
		newPost := Post{
			Text: text,
			Time: time.Now().Format("2 Jan 15:04"),
		}

		collection := Connect()
		_, err := collection.InsertOne(context.TODO(), newPost)
		if err != nil {
			log.Println(err)
		}
	}

	http.Redirect(w, req, "/", http.StatusFound)
}

//func deletePost(w http.ResponseWriter, req *http.Request) {}
