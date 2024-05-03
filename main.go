package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"goservice/Config"
	"goservice/Dto"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile("config.yml")

	var config = &Config.Configuration{}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Unable to decode config file, %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config file into configuration, %v", err)
	}
	fmt.Println("Database Cluster =\t", config.Database.Cluster)
	/*
		router := mux.NewRouter()
		router.HandleFunc("/", homeHandler).Methods("GET")
		router.HandleFunc("/api/v1/articles", articlesHandler).Methods("GET")

		http.ListenAndServe("0.0.0.0:8080", router)
	*/
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := Dto.NewServiceDto("hello World!")
	response.Metadata = []Dto.Hypermedia{
		Dto.Hypermedia{Relation: "next", Reference: "api/v1/articles"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	articles := []Dto.Article{Dto.Article{Id: 123, Title: "Hello!", Description: "Description of article", Content: "This is the content.  Enjoy it!"}}
	var hypermedia = make([]Dto.Hypermedia, len(articles))
	response := Dto.NewServiceDto(articles)
	for i, v := range articles {
		hypermedia[i].Reference = "Get"
		hypermedia[i].Relation = "api/v1/articles/" + strconv.FormatInt(int64(v.Id), 10)
	}
	response.Metadata = hypermedia
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
