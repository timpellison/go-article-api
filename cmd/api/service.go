package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"goservice/api"
	"goservice/config"
	"goservice/domain"
	"goservice/persistence"
	"os"
	"strconv"
	"strings"
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
	// let's get some config going!
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile("config.yml")

	var config = &config.Configuration{}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Unable to decode config file, %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config file into configuration, %v", err)
	}

	// now for some persistence
	dsn := "host=" + config.Database.Cluster +
		" user=" + config.Database.UserName +
		" port=" + strconv.FormatInt(int64(config.Database.Port), 10) +
		" dbname=" + config.Database.DatabaseName +
		" password=" + config.Database.Password +
		" sslmode=disable" +
		" TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Unable to connect to database, %v", err)
	}

	err = db.AutoMigrate(&domain.Article{})
	if err != nil {
		fmt.Printf("Unable to migrate database, %v", err)
	}

	// stand up our repository
	repo, err := persistence.NewArticleRepository(config, db)
	if err != nil {
		fmt.Printf("Unable to connect to database, %v", err)
	}

	server := api.NewServer(repo, config)
	server.Run()

	os.Exit(1)
}
