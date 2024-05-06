package Server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"goservice/Config"
	"goservice/Domain"
	"goservice/Dto"
	"goservice/Persistence"
	"net/http"
	"strconv"
)

type Server struct {
	repository    *Persistence.IArticleRepository
	configuration *Config.Configuration
}

type IArticleServer interface {
	Run()
	Cancel()
}

type ArticleServer struct {
	Configuration *Config.Configuration
	Repository    *Persistence.ArticleRepository
	Handler       *mux.Router
}

func NewServer(repository *Persistence.ArticleRepository, configuration *Config.Configuration) IArticleServer {
	router := mux.NewRouter()
	server := &ArticleServer{Configuration: configuration, Repository: repository, Handler: router}
	router.Handle("/api/v1/articles", newArticleHandler(server)).Methods("POST")
	router.HandleFunc("/", rootHandler).Methods("GET")
	router.Handle("/api/v1/articles", articlesHandler(server)).Methods("GET")
	router.Handle("/api/v1/articles/{id}", articleHandler(server)).Methods("GET")
	router.Handle("/api/v1/articles/{id}", deleteArticleHandler(server)).Methods("DELETE")

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
	dto := Dto.NewServiceDto("Articles API")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(dto)
	if err != nil {
		panic(err)
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
			article := &Dto.Article{Id: result.ID, Title: result.Title, Description: result.Description, Content: result.Content}
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
			var articleResult = make([]Dto.Article, len(*articles))
			var hypermedia = make([]Dto.Hypermedia, 0)

			for i, v := range *articles {
				articleResult[i].Id = v.ID
				articleResult[i].Title = v.Title
				articleResult[i].Description = v.Description
				articleResult[i].Content = v.Content
			}

			response := Dto.NewServiceDto(articleResult)
			response.Metadata = hypermedia
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
			var article Dto.Article
			err := json.NewDecoder(r.Body).Decode(&article)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Set("reason-phrase", err.Error())
				return
			}

			domainArticle := &Domain.Article{Title: article.Title, Description: article.Description, Content: article.Content}
			domainArticle = server.Repository.Add(domainArticle)

			article.Id = domainArticle.ID
			article.Title = domainArticle.Title
			article.Description = domainArticle.Description
			article.Content = domainArticle.Content
			response := Dto.NewServiceDto(article)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(response)
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
