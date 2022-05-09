package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/zelat/goharbor-client/apiv2"
	"github.com/zelat/goharbor-client/apiv2/config"
	"net/http"
)

const (
	username = "admin"
	password = "Harbor12345"
	apiURL   = "https://179.118.3.173/api"
)

func main() {
	ctx := context.Background()
	options := &config.Options{
		PageSize: 100,
		Timeout:  10,
		Sort:     "-name", // Sort all results in reversed alphabetical order
		Query:    "",
	}
	harborClient, err := apiv2.NewRESTClientForHost(apiURL, username, password, options)
	if err != nil {
		fmt.Print("连接harbor错误")
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	//列举所有Repositories信息
	repos, err := harborClient.ListRepositories(ctx, "secvector")
	var i int
	for i = 0; i < len(repos); i++ {
		fmt.Print(repos[i].Name, "\n")
		fmt.Print(repos[i].ProjectID, "\n")
		fmt.Print(repos[i].CreationTime, "\n")
		fmt.Print(repos[i].PullCount, "\n")
		fmt.Print(repos[i].Description, "\n")
	}
	//显示单个repository信息
	images, err := harborClient.GetRepository(ctx, "secvector", "api")
	fmt.Print("ImageName ", images.Name)
	fmt.Print("ImageDescription ", images.Description)
	//更新repository信息
	//images.Description = "test"
	//harborClient.UpdateRepository(ctx, "secvector", "api", images)
	//ListArtifact

}
