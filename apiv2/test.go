package apiv2

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/zelat/goharbor-client/v4/apiv2/config"
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
	harborClient, err := NewRESTClientForHost(apiURL, username, password, options)
	if err != nil {
		fmt.Print("连接harbor错误")
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	repos, err := harborClient.ListRepositories(ctx, "secvector")
	fmt.Print(repos)
}
