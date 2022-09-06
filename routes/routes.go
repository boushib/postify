package routes

import (
	"log"
	"net/http"

	"github.com/boushib/postify/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func InitRouter() {
	envMap, _ := godotenv.Read(".env")
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/api/v1/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/api/v1/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/api/v1/posts", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/v1/posts", controllers.DeletePost).Methods("DELETE")
	println("Server running on port " + envMap["PORT"] + "...")
	log.Fatal(http.ListenAndServe(":"+envMap["PORT"], router))
}
