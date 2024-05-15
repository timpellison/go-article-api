/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"goservice/config"
	"goservice/persistence"
	"os"
)

var dbName string
var databaseHost string
var databasePort int
var databaseUserName string
var databasePassword string

var conf *config.Configuration
var repository *persistence.IArticleRepository

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
	err := rootCmd.Execute()
	_ = fmt.Sprintf("The database name is %s", dbName)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbName, "dbname", "", fmt.Sprintf("When specified, this value is used for the datbaase name.  You can also set environment variable %s", persistence.DatabaseName))
	rootCmd.PersistentFlags().StringVar(&databaseHost, "dbhost", "", fmt.Sprintf("The hostname of the database to connect to.  You can also set environment variable %s", persistence.DatabaseHost))
	rootCmd.PersistentFlags().StringVar(&databasePassword, "dbpassword", "", fmt.Sprintf("The password to use for connecting to database.  You can also set the environment variable %s", persistence.DatabasePassword))
	rootCmd.PersistentFlags().IntVar(&databasePort, "dbport", 5432, fmt.Sprintf("The port # used to connect to the database.  You can also set the environment variable %s", persistence.DatabasePort))
	rootCmd.PersistentFlags().StringVar(&databaseUserName, "dbusername", "", fmt.Sprintf("The username to use to connect to the database.  You can also set the environment variable %s", persistence.DatabaseUserName))
}

func initRepository() {
	if os.Getenv(persistence.DatabasePassword) != "" {
		databasePassword = os.Getenv(persistence.DatabasePassword)
	}
}
