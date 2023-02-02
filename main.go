package main

import (
	"log"
	"net/http"

	"github.com/net-http/controllers"
)

func main() {
	mux := http.NewServeMux()
	newblog := controllers.NewBlogStore()

	mux.HandleFunc("/blog/allblog", newblog.GetAllBlogs)
	mux.HandleFunc("/blog/oneblog/", newblog.GetOneBlog)
	mux.HandleFunc("/blog/create", newblog.CreateBLog)
	mux.HandleFunc("/blog/delete/", newblog.DeleteBlog)

	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	name := r.URL.Query().Get("name")
	// 	if name == "" {
	// 		name = "World"
	// 	}
	// 	fmt.Fprintf(w, "Hello, %s", name)
	// })

	log.Println("Listening on http://localhost:8080/")

	http.ListenAndServe(":8080", mux)

}
