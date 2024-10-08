// Package main provides a simple example of how to use the generated SDK.
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
	folders, httpResponse, err := client.FoldersAPI.FoldersGet(auth).Execute()
	if err != nil {
		log.Fatal("Error!", err, httpResponse)
	}
	fmt.Println("Folders:")
	for _, folder := range folders.Data {
		fmt.Printf("%s (%s)\n", folder.Name, folder.FolderType)
	}
}
