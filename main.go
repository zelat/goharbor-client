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
	images, err := harborClient.GetRepository(ctx, "secvector", "httpd")
	fmt.Print("ImageName ", images.Name, "\n")
	fmt.Print("ImageDescription ", images.Description, "\n")
	//更新repository信息
	//images.Description = "test"
	//harborClient.UpdateRepository(ctx, "secvector", "api", images)

	//ListArtifact
	artifacts, err := harborClient.ListArtifacts(ctx, "secvector", "httpd")
	var j int
	for j = 0; j < len(artifacts); j++ {
		fmt.Print("j = ", j, "\n")
		fmt.Print("ArtifactsID = ", artifacts[j].ID, "\n")
		fmt.Print("ArtifactsDigest = ", artifacts[j].Digest, "\n")
		fmt.Print("ArtifactsReference = ", artifacts[j].References, "\n")
		//Get Artifact
		artifact, err := harborClient.GetArtifact(ctx, "secvector", "httpd", artifacts[j].Digest)
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Print("ID is ", artifact.ID, "\n")
		fmt.Print("RepositoryID is ", artifact.RepositoryID, "\n")
		fmt.Print("AdditionLinks is ", artifact.AdditionLinks, "\n")

		//Get Addition
		//	addition, err := harborClient.GetAddition(ctx, "secvector", "api", artifacts[j].Digest, "vulnerabilities")
		//	fmt.Print("addition", addition, "\n")
		//
		//Get Vuln
		//	vuln, err := harborClient.GetVulnerabilitiesAddition(ctx, "secvector", "api", artifacts[j].Digest)
		//	fmt.Print("vuln is ", vuln, "\n")
	}

	//Start scan Task
	//fmt.Print("Start Scan Task: ", artifacts[0].Digest, "\n")
	//reqID, err := harborClient.ScanArtifact(ctx, "secvector", "httpd", artifacts[0].Digest)
	//fmt.Print("requestID is ", reqID, "\n")

	//Get ReportLog
	//reportlog, err := harborClient.GetReportLog(ctx, "secvector", "httpd", artifacts[0].Digest, "b84fe5bd-42e8-4641-9727-e16a339fe510")
	//fmt.Print("ReportLog is ", reportlog, "\n")

	//Get Addition
	//addition, err := harborClient.GetAddition(ctx, "secvector", "httpd", artifacts[0].Digest, "vulnerabilities")
	//fmt.Print("addition is ", addition, "\n")

	//Get Vuln
	vuln, err := harborClient.GetVulnerabilitiesAddition(ctx, "secvector", "httpd", artifacts[0].Digest)
	fmt.Print("vuln is ", vuln.ReportType.GeneratedAt)
}
