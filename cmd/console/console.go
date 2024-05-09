package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"goservice/config"
	"goservice/domain"
	"goservice/persistence"
	"strconv"
	"strings"
)

func main() {
	Execute()
}

func Execute() {
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
}
