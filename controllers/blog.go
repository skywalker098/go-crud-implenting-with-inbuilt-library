package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/net-http/utils"
)

type Blog struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Blogstore struct {
	Blogs []Blog
}

func NewBlogStore() *Blogstore {
	return &Blogstore{Blogs: []Blog{}}
}

// Get all blogs
func (t Blogstore) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	t.LoadFromJson()
	if !utils.CheckMethod(r.Method, utils.GET) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// if r.Method != http.MethodGet {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	return
	// }""
	//marshal t.blogs
	blogs, err := json.Marshal(t.Blogs)
	if err != nil {
		log.Fatal(err)
	}

	utils.CustomRepsonseWriter(w, http.StatusOK, blogs)

	// fmt.Fprintf(w, "Get all blogs")
	// log.Println("Hello from GetAllBlogs", r.Method)

}

// Get one blog

func (t Blogstore) GetOneBlog(w http.ResponseWriter, r *http.Request) {
	t.LoadFromJson()

	//GET method check
	if !utils.CheckMethod(r.Method, utils.GET) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// if r.Method != http.MethodGet {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	return
	// }
	//get id from url
	id, err := utils.GetUrlParmId(r)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid blog id")
		return
	}
	for _, blog := range t.Blogs {
		if blog.Id == id {
			blogJson, err := json.Marshal(blog)
			if err != nil {
				log.Fatal(err)
			}
			utils.CustomRepsonseWriter(w, http.StatusOK, blogJson)
			return
		}
	}
	utils.CustomRepsonseWriter(w, http.StatusNotFound, nil)
	// log.Printf("Geone blog with id %d", id)
	// log.Println("Hello from GetOneBlog", r.Method)

}

// Create a blog
func (t Blogstore) CreateBLog(w http.ResponseWriter, r *http.Request) {
	t.LoadFromJson()

	if !utils.CheckMethod(r.Method, utils.POST) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// if r.Method!= http.MethodPost {
	if !utils.CheckMethod(r.Method, utils.POST) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// if r.Method != http.MethodPost {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	return
	// }
	var blog Blog
	//decode request body into the struct
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	blog.Id = t.AddnewId()

	t.Blogs = append(t.Blogs, blog)
	t.SavetoJson()

	data, err := json.Marshal(blog)
	if err != nil {
		log.Fatal(err)
	}
	utils.CustomRepsonseWriter(w, http.StatusCreated, data)

	// log.Printf("Add blog %+v", blog)
	// fmt.Fprintf(w, blog.Body)
	// fmt.Println("Hello from CreateBLog", r.Method)

}

// Delete a blog
func (t Blogstore) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	t.LoadFromJson()

	if !utils.CheckMethod(r.Method, utils.DELETE) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// if r.Method != http.MethodDelete {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	return
	// }

	id, err := utils.GetUrlParmId(r)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	for i, blog := range t.Blogs {
		if blog.Id == id {
			t.Blogs = append(t.Blogs[:i], t.Blogs[i+1:]...)
			t.SavetoJson()
			return
		}
	}
	t.SavetoJson()

	utils.CustomRepsonseWriter(w, http.StatusNotFound, nil)

	// fmt.Fprintf(w, "Delete blog %d", id)
	// log.Printf("Blog deleted with id %d", id)
	// fmt.Println("Hello from DeleteBLog", r.Method)

}
