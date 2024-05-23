/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"goservice/config"
	"goservice/persistence"
	"log"
	"os"
	"strconv"
)

var dbNameFlag string
var dbPasswordFlag string
var dbHostFlag string
var dbUserNameFlag string
var dbPortFlag int

var logger *log.Logger
var repository *persistence.ArticleRepository

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dunectl",
	Short: "Data importer for articles.",
	Long: `The articles http-server is built to provide CRUD as well as search operations for the Dune Library.

`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbNameFlag, "dbname", "", fmt.Sprintf("When specified, this value is used for the datbaase name.  You can also set environment variable %s", persistence.DatabaseName))
	rootCmd.PersistentFlags().StringVar(&dbHostFlag, "dbhost", "", fmt.Sprintf("The hostname of the database to connect to.  You can also set environment variable %s", persistence.DatabaseCluster))
	rootCmd.PersistentFlags().StringVar(&dbPasswordFlag, "dbpassword", "", fmt.Sprintf("The password to use for connecting to database.  You can also set the environment variable %s", persistence.DatabasePassword))
	rootCmd.PersistentFlags().IntVar(&dbPortFlag, "dbport", 5432, fmt.Sprintf("The port # used to connect to the database.  You can also set the environment variable %s", persistence.DatabasePort))
	rootCmd.PersistentFlags().StringVar(&dbUserNameFlag, "dbusername", "", fmt.Sprintf("The username to use to connect to the database.  You can also set the environment variable %s", persistence.DatabaseUserName))
}

func getConfiguration() (*config.Configuration, error) {
	var dbName, dbHost, dbUserName, dbPassword string
	var dbPort int

	name, noname := os.LookupEnv(persistence.DatabaseName)
	host, nohost := os.LookupEnv(persistence.DatabaseCluster)
	user, nouser := os.LookupEnv(persistence.DatabaseUserName)
	password, nopassword := os.LookupEnv(persistence.DatabasePassword)

	port, noport := os.LookupEnv(persistence.DatabasePort)

	if dbNameFlag == "" && !noname {
		logger.Fatal("database name is required")
	} else {
		if dbNameFlag != "" {
			dbName = dbNameFlag
		} else {
			dbName = name
		}
	}

	if dbHostFlag == "" && !nohost {
		logger.Fatal("database host is required")
	} else {
		if dbHostFlag != "" {
			dbHost = dbHostFlag
		} else {
			dbHost = host
		}
	}

	if dbUserNameFlag == "" && !nouser {
		logger.Fatal("database username is required")
	} else {
		if dbUserNameFlag != "" {
			dbUserName = dbUserNameFlag
		} else {
			dbUserName = user
		}
	}

	if dbPasswordFlag == "" && !nopassword {
		logger.Fatal("database password is required")
	} else {
		if dbPasswordFlag != "" {
			dbPassword = dbPasswordFlag
		} else {
			dbPassword = password
		}
	}

	if dbPortFlag == 0 && !noport {
		logger.Fatal("database port is required")
	} else {
		if dbPortFlag != 0 {
			dbPort = dbPortFlag
		} else {
			p, err := strconv.Atoi(port)
			if err != nil {
				logger.Fatal("database port is not a number")
			}
			dbPort = p
		}
	}
	database := &config.DatabaseConfiguration{Cluster: dbHost, Port: dbPort, DatabaseName: dbName, UserName: dbUserName, Password: dbPassword}
	conf := &config.Configuration{Database: *database}
	return conf, nil
}
