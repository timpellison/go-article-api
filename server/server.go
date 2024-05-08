package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"goservice/config"
	"goservice/domain"
	"goservice/dto"
	"goservice/persistence"
	"net/http"
	"strconv"
)

type Server struct {
	repository    *persistence.IArticleRepository
	configuration *config.Configuration
}

type IArticleServer interface {
	Run()
	Cancel()
}

type ArticleServer struct {
	Configuration *config.Configuration
	Repository    *persistence.ArticleRepository
	Handler       *mux.Router
}

func NewServer(repository *persistence.ArticleRepository, configuration *config.Configuration) IArticleServer {
	router := mux.NewRouter()
	server := &ArticleServer{Configuration: configuration, Repository: repository, Handler: router}
	router.Handle("/api/v1/articles", newArticleHandler(server)).Methods("POST")
	router.HandleFunc("/", rootHandler).Methods("GET")
	router.Handle("/api/v1/articles", articlesHandler(server)).Methods("GET")
	router.Handle("/api/v1/articles/{id}", articleHandler(server)).Methods("GET")
	router.Handle("/api/v1/articles/{id}", deleteArticleHandler(server)).Methods("DELETE")
	router.StrictSlash(false)
	return server
}

func (server *ArticleServer) Run() {
	var serverPort = "0.0.0.0:" + strconv.Itoa(server.Configuration.Server.Port)
	err := http.ListenAndServe(serverPort, server.Handler)
	if err != nil {
		return
	}
	defer server.Cancel()
}

func (server *ArticleServer) Cancel() {

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "%s", "Articles API")
	if err != nil {
		return
	}
}

func articleHandler(server *ArticleServer) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			pathId := mux.Vars(r)["id"]
			id, err := strconv.Atoi(pathId)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			result := server.Repository.GetOne(uint(id))
			if result == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			article := &dto.Article{Id: result.ID, Title: result.Title, Description: result.Description, Content: result.Content}
			var media = []dto.Hypermedia{
				dto.Hypermedia{Relation: "delete", Reference: "/api/v1/articles/" + strconv.FormatInt(int64(id), 10)},
			}
			article.Metadata = media

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(article)
			if err != nil {
				w.Header().Set("reason-phrase", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		})
}

func articlesHandler(server *ArticleServer) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			articles := server.Repository.GetMany()
			var articleResult = make([]dto.Article, len(*articles))

			for i, v := range *articles {
				articleResult[i].Id = v.ID
				articleResult[i].Title = v.Title
				articleResult[i].Description = v.Description
				articleResult[i].Content = v.Content

				var hypermedia = []dto.Hypermedia{
					dto.NewHypermedia("delete", "/api/v1/articles/"+strconv.FormatInt(int64(v.ID), 10)),
				}
				articleResult[i].Metadata = hypermedia
			}

			response := articleResult

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		})
}

func newArticleHandler(server *ArticleServer) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var article dto.Article
			err := json.NewDecoder(r.Body).Decode(&article)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Set("reason-phrase", err.Error())
				return
			}

			domainArticle := &domain.Article{Title: article.Title, Description: article.Description, Content: article.Content}
			domainArticle = server.Repository.Add(domainArticle)

			article.Id = domainArticle.ID
			article.Title = domainArticle.Title
			article.Description = domainArticle.Description
			article.Content = domainArticle.Content

			var hypermedia = []dto.Hypermedia{
				dto.NewHypermedia("delete", "/api/v1/articles/"+strconv.FormatInt(int64(article.Id), 10)),
			}
			article.Metadata = hypermedia

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(article)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		})
}

func deleteArticleHandler(server *ArticleServer) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.Atoi(mux.Vars(request)["id"])
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Header().Set("reason-phrase", err.Error())
		}
		article := server.Repository.GetOne(uint(id))
		fmt.Printf("Removing article %v", article)
		if article != nil {
			server.Repository.Delete(article)
		}
	})
}
