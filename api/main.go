package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"goservice/config"
	_ "goservice/docs" // docs is generated by Swag CLI, you have to import it.
	"goservice/domain"
	"goservice/dto"
	"goservice/persistence"
	"log"
	"net/http"
	"os"
	"strconv"
)

type IArticleServer interface {
	Run()
}

type ArticleServer struct {
	Configuration *config.Configuration
	Repository    *persistence.ArticleRepository
	Logger        *log.Logger
	Handler       *mux.Router
}

func NewServer(repository *persistence.ArticleRepository, configuration *config.Configuration) IArticleServer {
	router := mux.NewRouter()
	server := &ArticleServer{Configuration: configuration, Repository: repository, Handler: router}
	router.Handle("/api/v1/articles", newArticleHandler(server)).Methods("POST")
	router.Handle("/api/v1/articles", articlesHandler(server)).Methods("GET")
	router.Handle("/api/v1/articles/{id}", articleHandler(server)).Methods("GET")
	router.Handle("/api/v1/articles/{id}", deleteArticleHandler(server)).Methods("DELETE")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.StrictSlash(false)

	log.SetOutput(os.Stdout)
	server.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	return server
}

func (server *ArticleServer) Run() {
	var serverPort = ":" + strconv.Itoa(server.Configuration.Server.Port)
	err := http.ListenAndServe(serverPort, server.Handler)
	if err != nil {
		return
	}
	defer server.Cancel()
}

func (server *ArticleServer) Cancel() {

}

// Articles godoc
//
//		@Summary		Returns an article by its ID
//		@Description	Returns an article by its ID
//		@Tags			articles
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	dto.Article
//		@Failure		500
//	    @Param          Id path uint true "Article Id"
//		@Router			/api/v1/articles/{Id} [get]
func articleHandler(server *ArticleServer) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			pathId := mux.Vars(r)["id"]
			id, err := strconv.Atoi(pathId)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			result, err := server.Repository.GetOne(uint(id))
			if result == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			article := &dto.Article{Id: result.ID, Title: result.Title, Description: result.Description, Content: result.Content}
			var media = []dto.Hypermedia{
				{Relation: "delete", Reference: "/api/v1/articles/" + strconv.FormatInt(int64(id), 10)},
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

// Articles godoc
//
//	@Summary		Returns all articles
//	@Description	Returns all articles
//	@Tags			articles
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dto.Article
//	@Failure		500
//	@Router			/api/v1/articles [get]
func articlesHandler(server *ArticleServer) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			articles, err := server.Repository.GetMany()
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			var articleResult = make([]dto.Article, len(*articles))
			server.Logger.Printf("Total length of articles is %d\n", len(*articles))

			for i, v := range *articles {
				var hypermedia = []dto.Hypermedia{
					dto.NewHypermedia("delete", "/api/v1/articles/"+strconv.FormatInt(int64(v.ID), 10)),
					dto.NewHypermedia("get", "/api/v1/articles/"+strconv.FormatInt(int64(v.ID), 10)),
				}
				articleResult[i].Id = v.ID
				articleResult[i].Title = v.Title
				articleResult[i].Description = v.Description
				articleResult[i].Content = v.Content
				articleResult[i].Metadata = hypermedia
				server.Logger.Printf("Article added.  Id is %d, Title is %v\n", articleResult[i].Id, articleResult[i].Title)
			}

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(&articleResult)

			if err != nil {
				server.Logger.Printf("Error encoding articles to json: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		})
}

// Articles godoc
//
//		@Summary		Add a new Article to the system
//		@Description	Add a new Article to the system
//		@Tags			articles
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	dto.ArticleData
//		@Failure		500
//	    @Param          article body dto.ArticleData true "Article"
//		@Router			/api/v1/articles [post]
func newArticleHandler(server *ArticleServer) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var article dto.ArticleData
			err := json.NewDecoder(r.Body).Decode(&article)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Set("reason-phrase", err.Error())
				return
			}

			domainArticle := &domain.Article{Title: article.Title, Description: article.Description, Content: article.Content}
			domainArticle, err = server.Repository.Add(domainArticle)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			newArticle := &dto.Article{}

			newArticle.Id = domainArticle.ID
			newArticle.Title = domainArticle.Title
			newArticle.Description = domainArticle.Description
			newArticle.Content = domainArticle.Content

			var hypermedia = []dto.Hypermedia{
				dto.NewHypermedia("delete", "/api/v1/articles/"+strconv.FormatInt(int64(newArticle.Id), 10)),
			}
			newArticle.Metadata = hypermedia

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(newArticle)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		})
}

// Articles godoc
//
//			@Summary		Delete the article identified by Id
//			@Description	Delete the article identified by Id
//			@Tags			articles
//			@Accept			json
//			@Produce		json
//			@Success		200
//			@Failure		500
//	     	@Failure        400
//		    @Param          Id path uint true "Article Id"
//			@Router			/api/v1/articles/{Id} [delete]
func deleteArticleHandler(server *ArticleServer) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.Atoi(mux.Vars(request)["id"])
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Header().Set("reason-phrase", err.Error())
		}
		article, err := server.Repository.GetOne(uint(id))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Header().Set("reason-phrase", err.Error())
		}
		fmt.Printf("Removing article %v", article)
		if article != nil {
			server.Repository.Delete(article)
		}
	})
}