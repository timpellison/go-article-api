package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"goservice/config"
	"goservice/domain"
	"goservice/persistence"
	"goservice/server"
	"strconv"
	"strings"
)

/*
func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()


	return nil
}
*/

func main() {
	//ctx := context.Background()
	//if err := run(ctx, os.Stdout, os.Args[1:]); err != nil {
	//	fmt.Fprintf(os.Stderr, "%s\n", err)
	//	os.Exit(1)
	//}

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
	repo := persistence.NewArticleRepository(config, db)

	server := server.NewServer(repo, config)
	server.Run()
}
