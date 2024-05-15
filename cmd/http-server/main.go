package main

import (
	"fmt"
	"goservice/api"
	"goservice/config"
	_ "goservice/docs"
	"goservice/persistence"
	"os"
)

// @title Articles API
// @version 1.0
// @description This is the articles API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	var conf = config.NewConfiguration("./config.yml")

	// stand up our repository
	repo, err := persistence.NewArticleRepository(conf)
	if err != nil {
		fmt.Printf("Unable to connect to database, %v", err)
	}

	server := api.NewServer(repo, conf)
	server.Run()

	os.Exit(1)
}
