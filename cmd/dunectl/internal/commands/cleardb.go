package commands

import (
	"github.com/spf13/cobra"
	"goservice/persistence"
)

var clearDatabaseCmd = &cobra.Command{
	Use:   "cleardb",
	Short: "Remove all items from database",
	Run: func(cmd *cobra.Command, args []string) {

		config, err := getConfiguration()
		if err != nil {
			logger.Fatal(err)
		}
		repo, err := persistence.NewArticleRepository(config)
		if err != nil {
			logger.Fatal(err)
		}

		articles, err := repo.GetMany()
		if err != nil {
			logger.Fatal(err)
		}

		for _, v := range *articles {
			err = repo.Delete(&v)
			if err != nil {
				logger.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(clearDatabaseCmd)
}
