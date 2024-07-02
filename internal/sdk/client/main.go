package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

func main() {
	apiKey := os.Getenv("RETOOL_ACCESS_TOKEN")

	if apiKey == "" {
		fmt.Println("Please set the RETOOL_ACCESS_TOKEN environment variable")
		return
	}

	config := api.NewConfiguration()
	config.Host = "localhost:3000"
	config.Scheme = "http"
	config.Servers = api.ServerConfigurations{api.ServerConfiguration{
		URL: "/api/v2",
	},
	}
	client := api.NewAPIClient(config)
	auth := context.WithValue(context.Background(), api.ContextAccessToken, apiKey)
	folders, httpResponse, error := client.FoldersAPI.FoldersGet(auth).Execute()
	if error != nil {
		log.Fatal("Error!", error, httpResponse)
	} else {
		fmt.Println("Folders:")
		for _, folder := range folders.Data {
			fmt.Printf("%s (%s)\n", folder.Name, folder.FolderType)
		}
	}
}
