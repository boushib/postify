package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/boushib/go-blog/config"
	"github.com/boushib/go-blog/models"
	"github.com/gorilla/mux"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []models.Post
	query := `SELECT id, title, content FROM posts`
	rows, err := config.DB.Query(query)

	if err != nil {
		fmt.Println("Error querying database", err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var title string
		var content string

		rows.Scan(&id, &title, &content)
		post := models.Post{
			Id:      id,
			Title:   title,
			Content: content,
		}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := `SELECT id, title, content FROM posts WHERE id = $1`
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		fmt.Println("Error parsing id")
		return
	}

	row, err := config.DB.Query(query, id)

	defer row.Close()

	if row.Next() {
		var id uint64
		var title string
		var content string
		row.Scan(&id, &title, &content)
		post := models.Post{
			Id:      id,
			Title:   title,
			Content: content,
		}
		json.NewEncoder(w).Encode(post)
	}

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	query := `INSERT INTO posts (title, content) VALUES ($1, $2)`
	_, err := config.DB.Exec(query, post.Title, post.Content)

	if err != nil {
		fmt.Println("Error inserting post into database")
		return
	}

	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	query := `UPDATE posts SET title = $1, content = $2 WHERE id = $3`
	_, err := config.DB.Exec(query, post.Title, post.Content, post.Id)

	if err != nil {
		fmt.Println("Error updating post")
		return
	}

	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := `DELETE FROM posts WHERE id = $1`
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		fmt.Println("Error parsing id")
		return
	}

	_, err = config.DB.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting post")
		return
	}

	json.NewEncoder(w).Encode("Post deleted successfully")
}
