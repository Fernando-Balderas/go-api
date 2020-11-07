package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      string `json:"ID"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "home called"}`))
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Articles)
}

func getArticleByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	key := vars["articleId"]
	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	vars := mux.Vars(r)
	id := vars["articleId"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var modifiedArticle Article
	json.Unmarshal(reqBody, &modifiedArticle)

	for idx, article := range Articles {
		if article.ID == id {
			Articles[idx] = modifiedArticle
		}
	}

	w.Write([]byte(`{"message": "put called"}`))
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["articleId"]

	for idx, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:idx], Articles[idx+1:]...)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "article deleted"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func handleRequests() {
	rtr := mux.NewRouter().StrictSlash(true)
	rtr.HandleFunc("/", home).Methods(http.MethodGet)
	rtr.HandleFunc("/articles", getArticles).Methods(http.MethodGet)
	rtr.HandleFunc("/articles/{articleId}", getArticleByID).Methods(http.MethodGet)
	rtr.HandleFunc("/articles", createNewArticle).Methods(http.MethodPost)
	rtr.HandleFunc("/articles/{articleId}", updateArticle).Methods(http.MethodPut)
	rtr.HandleFunc("/articles/{articleId}", deleteArticle).Methods(http.MethodDelete)
	rtr.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8080", rtr))
}

func main() {
	fmt.Println("Rest API with Mux Routers")
	Articles = []Article{
		Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
