package commands

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"goservice/domain"
	"goservice/persistence"
	"os"
	"time"
)

type duneArticle struct {
	VolumeInfo  volumeInfo
	Description string `json:"description"`
}

type volumeInfo struct {
	Title       string   `json:"title"`
	Publisher   string   `json:"publisher"`
	PublishDate string   `json:"publishedDate"`
	Authors     []string `json:"authors"`
}

var appendDatabaseCmd = &cobra.Command{
	Use:   "append",
	Short: "Add articles to database from json file",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := getConfiguration()
		if err != nil {
			logger.Fatal(err)
		}
		repo, err := persistence.NewArticleRepository(config)
		if err != nil {
			logger.Fatal(err)
		}

		if len(args) == 0 {
			logger.Fatal("You must pass the path and file name of the json file to import.")
		}

		fileName := args[0]
		var allData []duneArticle
		plan, err := os.ReadFile(fileName)
		if err != nil {
			logger.Fatal(err)
		}

		err = json.Unmarshal(plan, &allData)
		if err != nil {
			logger.Fatal(err)
		}

		var publishDate time.Time

		for _, article := range allData {
			description := fmt.Sprintf("Published by %s on %s by author(s) %s", article.VolumeInfo.Publisher, article.VolumeInfo.PublishDate, article.VolumeInfo.Authors)
			if len(article.VolumeInfo.PublishDate) > 4 {
				publishDate, _ = time.Parse("2006-01-02", article.VolumeInfo.PublishDate)
			}

			fmt.Println(publishDate)
			model := domain.Article{Title: article.VolumeInfo.Title, Description: description, Content: article.Description, PublishDate: publishDate}
			_, err = repo.Add(&model)
			if err != nil {
				logger.Fatal(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(appendDatabaseCmd)
}
