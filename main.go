package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User models
type User struct{
	ID string `json:"id"`
	Name string `json:"Name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

// Post Model
type Post struct {
	ID  string `json:"id"`
	Caption  string `json:"caption"`
	ImageUrl string `json:"imageurl`
	Timestamp string `json:"timestamp"`
	User *User `json:"user"` // user uploads the post
}

// init users var as a slice User struct
var users []User
var posts []Post

// create a new user
func createUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application.json")
	var book User
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000000)) //mock id
	users = append(users, book)
	json.NewEncoder(w).Encode(&User{})
}

// to get single user through id
func getUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application.json")
	params := mux.Vars(r)
	// loop thorugh users
	for _, item:= range users{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

// create post
func createPost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application.json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000000)) //mock id
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&Post{})
}

// get post
func getPost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application.json")
	params := mux.Vars(r)
	// loop thorugh posts
	for _, item:= range posts{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

// find user through post
func findAllPost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application.json")
	params := mux.Vars(r)
	for _, pst:= range posts{
		if pst.ID == params["id"]{
			for _, usid:= range users{
				if usid.ID == params["id"]{
					json.NewEncoder(w).Encode(item)
					return
				}
			}
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func main(){

	r := mux.NewRouter()

	// route handlers
	r.HandleFunc("/api/users", createUser).Method("POST") // create user
	r.HandleFunc("/api/users/{id}", getUser).Method("GET") // get user through ID
	r.HandleFunc("/api/posts",createPost).method("POST") // creating an post
	r.HandleFunc("/api/posts/{id}", getPost).method("GET") // geting post from users
	r.HandleFunc("/api/posts/user/{id}", findAllPost).method("GET") // list of all the post user made 

	// Port
	if err := http.ListenAndServe(":8080",r);
	err !=nil{
		log.Fatal(err)
	}
}



